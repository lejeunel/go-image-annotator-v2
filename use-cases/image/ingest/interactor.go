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
)

var errCtx = "ingesting image"

type Interactor struct {
	hasher         Hasher
	imageRepo      ImageRepo
	collectionRepo CollectionRepo
	annotationRepo AnnotationRepo
	labelRepo      LabelRepo
	artefactRepo   ast.ArtefactRepo
	imageDecoder   ImageDecoder
}

func NewInteractor(imageRepo ImageRepo, collectionRepo CollectionRepo,
	labelRepo LabelRepo, annotationRepo AnnotationRepo,
	artefactRepo ast.ArtefactRepo, hasher Hasher, decoder ImageDecoder) *Interactor {
	return &Interactor{imageRepo: imageRepo, collectionRepo: collectionRepo,
		annotationRepo: annotationRepo, labelRepo: labelRepo,
		artefactRepo: artefactRepo, hasher: hasher, imageDecoder: decoder}
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	collection, ok := i.findCollectionByName(r.Collection, out)
	if !ok {
		return
	}

	imageId := im.NewImageId()
	image, ok := i.createImage(imageId, *collection, r.Labels, r.BoundingBoxes, out)
	if !ok {
		return
	}

	data, _, err := i.imageDecoder.Decode(r.Data)
	if err != nil {
		out.ErrValidation(fmt.Errorf("%v: reading image data: %w", errCtx, e.ErrValidation))
		return
	}

	hash := i.hasher.Hash(data)

	ok = i.ingestRawData(imageId, data, hash, out)
	if !ok {
		return
	}

	ok = i.ingestImage(image, hash, out)
	if !ok {
		i.imageRepo.Delete(image.Id)
		i.artefactRepo.Delete(image.Id)
		return
	}

	out.Success(NewIngestionResponse(image))

}

func (i *Interactor) ingestRawData(id im.ImageId, data []byte, hash string, out OutputPort) bool {

	if i.duplicateImageExists(hash, out) {
		return false
	}

	if err := i.artefactRepo.Store(id, data); err != nil {
		out.ErrInternal(fmt.Errorf("%v: storing image data in artefact repository: %w", errCtx, e.ErrInternal))
		return false
	}

	return true

}

func (i *Interactor) createImage(id im.ImageId, collection clc.Collection, labelNames []string,
	bboxes []BoundingBoxRequest, out OutputPort) (*im.Image, bool) {
	image := im.NewImage(id, collection)

	if ok := i.appendLabels(image, labelNames, out); !ok {
		return nil, false
	}

	if ok := i.appendBoundingBoxes(image, bboxes, out); !ok {
		return nil, false
	}

	return image, true

}

func (i *Interactor) appendLabels(image *im.Image, labelNames []string, out OutputPort) bool {
	for _, labelName := range labelNames {
		label, ok := i.findLabelByName(labelName, out)
		if !ok {
			return false
		}
		if err := image.AddLabel(label); err != nil {
			out.ErrValidation(fmt.Errorf("%v: %w", errCtx, err))
			return false
		}
	}
	return true

}
func (i *Interactor) appendBoundingBoxes(image *im.Image, bboxes []BoundingBoxRequest, out OutputPort) bool {
	for _, bbox := range bboxes {
		label, ok := i.findLabelByName(bbox.Label, out)
		if !ok {
			return false
		}
		box_ := a.NewBoundingBox(a.NewAnnotationId(), bbox.Xc, bbox.Yc, bbox.Width, bbox.Height, *label)
		if err := image.AddBoundingBox(*box_); err != nil {
			out.ErrValidation(fmt.Errorf("%v: %w", errCtx, err))
			return false
		}
	}
	return true

}

func (i *Interactor) ingestImage(image *im.Image, hash string, out OutputPort) bool {
	if err := i.imageRepo.AddImage(image.Id, hash); err != nil {
		out.ErrInternal(fmt.Errorf("%v: adding image: %w", errCtx, e.ErrInternal))
		return false
	}

	if err := i.imageRepo.AddImageToCollection(image.Id, image.Collection.Id); err != nil {
		out.ErrInternal(fmt.Errorf("%v: adding image to collection %v: %w", errCtx, image.Collection.Name, e.ErrInternal))
		return false
	}

	for _, label := range image.Labels {
		if err := i.annotationRepo.AddImageLabel(an.NewAnnotationId(), image.Id, image.Collection.Id, label.Label.Id); err != nil {
			out.ErrInternal(fmt.Errorf("%v: ingesting label: %w", errCtx, e.ErrInternal))
			return false
		}
	}

	for _, box := range image.BoundingBoxes {
		if err := i.annotationRepo.AddBoundingBox(image.Id, image.Collection.Id, *box); err != nil {
			out.ErrInternal(fmt.Errorf("%v: ingesting bounding box: %w", errCtx, e.ErrInternal))
			return false
		}
	}
	return true

}

func (i *Interactor) findCollectionByName(name string, out OutputPort) (*clc.Collection, bool) {
	collection, err := i.collectionRepo.FindCollectionByName(name)
	baseErrMsg := fmt.Sprintf("%v: finding collection with name %v", errCtx, name)
	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
		return nil, false
	case errors.Is(err, e.ErrInternal):
		out.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return nil, false
	}
	return collection, true

}

func (i *Interactor) findLabelByName(name string, out OutputPort) (*lbl.Label, bool) {
	baseErrMsg := fmt.Sprintf("%v: fetching label %v", errCtx, name)
	label, err := i.labelRepo.FindLabelByName(name)
	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
		return nil, false
	case errors.Is(err, e.ErrInternal):
		out.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return nil, false
	}
	return label, true

}

func (i *Interactor) duplicateImageExists(hash string, out OutputPort) bool {

	errCtx_ := fmt.Sprintf("%v: searching for duplicate image using hash", errCtx)
	duplicateId, err := i.imageRepo.FindImageIdByHash(hash)
	switch {
	case errors.Is(e.ErrNotFound, err):
		return false
	case duplicateId != nil:
		out.ErrDuplication(fmt.Errorf("%v: found duplicate with id: %v: %w", errCtx_, duplicateId, e.ErrValidation))
		return true
	default:
		out.ErrInternal(fmt.Errorf("%v: %w", errCtx_, e.ErrInternal))
		return true
	}
}
