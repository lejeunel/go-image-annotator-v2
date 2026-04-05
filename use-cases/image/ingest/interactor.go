package ingest

import (
	"errors"
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"

	ast "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	an "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	"log/slog"
)

type Interactor struct {
	hasher         Hasher
	imageRepo      ImageRepo
	collectionRepo CollectionRepo
	annotationRepo AnnotationRepo
	labelRepo      LabelRepo
	artefactRepo   ast.ArtefactRepo
	imageDecoder   ImageDecoder
	logger         *slog.Logger
}

func NewInteractor(imageRepo ImageRepo, collectionRepo CollectionRepo,
	labelRepo LabelRepo, annotationRepo AnnotationRepo,
	artefactRepo ast.ArtefactRepo, hasher Hasher, decoder ImageDecoder) *Interactor {
	return &Interactor{imageRepo: imageRepo, collectionRepo: collectionRepo,
		annotationRepo: annotationRepo, labelRepo: labelRepo,
		artefactRepo: artefactRepo, hasher: hasher, imageDecoder: decoder,
		logger: logging.NewNoOpLogger(),
	}
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	collection, err := i.findCollectionByName(r.Collection)
	if err != nil {
		i.handleError(err, out)
		return
	}

	imageId := im.NewImageId()
	image, err := i.createImage(imageId, *collection, r.Labels, r.BoundingBoxes)
	if err != nil {
		i.handleError(err, out)
		return
	}

	data, _, err := i.imageDecoder.Decode(r.Data)
	if err != nil {
		i.handleError(err, out)
		return
	}

	hash := i.hasher.Hash(data)
	if err := i.ingestRawData(imageId, data, hash); err != nil {
		i.handleError(err, out)
		return

	}

	if err := i.ingestImage(image, hash); err != nil {
		i.handleError(err, out)
		i.imageRepo.Delete(image.Id)
		i.artefactRepo.Delete(image.Id)
		return
	}

	out.Success(NewIngestionResponse(image))

}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "ingesting image"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)

	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(err)
	case errors.Is(err, e.ErrDuplicate):
		out.ErrDuplication(err)
	case errors.Is(err, e.ErrValidation):
		out.ErrValidation(err)
	default:
		out.ErrInternal(err)
	}
}
func (i *Interactor) ingestRawData(id im.ImageId, data []byte, hash string) error {

	if err := i.ensureDuplicateImageDoesNotExists(hash); err != nil {
		return err
	}

	if err := i.artefactRepo.Store(id, data); err != nil {
		return err
	}

	return nil

}

func (i *Interactor) createImage(id im.ImageId, collection clc.Collection, labelNames []string,
	bboxes []BoundingBoxRequest) (*im.Image, error) {
	image := im.NewImage(id, collection)

	if err := i.appendLabels(image, labelNames); err != nil {
		return nil, err
	}

	if err := i.appendBoundingBoxes(image, bboxes); err != nil {
		return nil, err
	}

	return image, nil

}

func (i *Interactor) appendLabels(image *im.Image, labelNames []string) error {
	for _, labelName := range labelNames {
		label, err := i.findLabelByName(labelName)
		if err != nil {
			return err
		}
		if err := image.AddLabel(label); err != nil {
			return err
		}
	}
	return nil

}
func (i *Interactor) appendBoundingBoxes(image *im.Image, bboxes []BoundingBoxRequest) error {
	baseErr := fmt.Errorf("appending bounding boxes")
	for _, bbox := range bboxes {
		label, err := i.findLabelByName(bbox.Label)
		if err != nil {
			return fmt.Errorf("%w: %w", baseErr, err)
		}
		box_ := a.NewBoundingBox(a.NewAnnotationId(), bbox.Xc, bbox.Yc, bbox.Width, bbox.Height, *label)
		if err := image.AddBoundingBox(*box_); err != nil {
			return fmt.Errorf("%w: %w", baseErr, err)
		}
	}
	return nil

}

func (i *Interactor) ingestImage(image *im.Image, hash string) error {

	if err := i.imageRepo.AddImage(image.Id, hash); err != nil {
		return fmt.Errorf("adding image: %w", err)
	}

	if err := i.imageRepo.AddImageToCollection(image.Id, image.Collection.Id); err != nil {
		return fmt.Errorf("adding image to collection: %w", err)
	}

	for _, label := range image.Labels {
		if err := i.annotationRepo.AddImageLabel(an.NewAnnotationId(), image.Id, image.Collection.Id, label.Label.Id); err != nil {
			return fmt.Errorf("adding image label to collection: %w", err)
		}
	}

	for _, box := range image.BoundingBoxes {
		if err := i.annotationRepo.AddBoundingBox(image.Id, image.Collection.Id, *box); err != nil {
			return fmt.Errorf("adding bounding box: %w", err)
		}
	}
	return nil

}

func (i *Interactor) findCollectionByName(name string) (*clc.Collection, error) {
	collection, err := i.collectionRepo.FindCollectionByName(name)
	baseErr := fmt.Errorf("finding collection with name %v", name)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", baseErr, err)
	}
	return collection, nil

}

func (i *Interactor) findLabelByName(name string) (*lbl.Label, error) {
	baseErr := fmt.Errorf("fetching label by name %v", name)
	label, err := i.labelRepo.FindLabelByName(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", baseErr, err)
	}
	return label, nil

}

func (i *Interactor) ensureDuplicateImageDoesNotExists(hash string) error {

	baseErr := fmt.Errorf("ensuring that duplicate image does not exist using hash")
	duplicateId, err := i.imageRepo.FindImageIdByHash(hash)
	if duplicateId != nil {
		return fmt.Errorf("%w: found duplicate image with id %v: %w", baseErr, *duplicateId, e.ErrDuplicate)

	}

	if errors.Is(err, e.ErrNotFound) {
		return nil
	}
	return err
}
