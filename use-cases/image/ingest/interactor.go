package ingest

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"

	e "github.com/lejeunel/go-image-annotator-v2/errors"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

var errCtx = "ingesting image"

type Interactor struct {
	repo         Repo
	artefactRepo im.ArtefactRepo
}

func NewInteractor(repo Repo, artefactRepo im.ArtefactRepo) *Interactor {
	return &Interactor{repo: repo, artefactRepo: artefactRepo}
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	collection, ok := i.findCollectionByName(r.Collection, out)
	if !ok {
		return
	}

	imageId := im.NewImageId()
	ok = i.ingestRawData(imageId, r.Reader, out)
	if !ok {
		return
	}

	image, ok := i.createImage(imageId, *collection, r.Labels, r.BoundingBoxes, out)
	if !ok {
		return
	}

	ok = i.ingestImage(image, out)
	if !ok {
		i.repo.DeleteImage(image.Id)
		i.artefactRepo.Delete(image.Id)
		return
	}

	out.Success(NewIngestionResponse(image))

}

func (i *Interactor) ingestRawData(id im.ImageId, reader im.ImageReader, out OutputPort) bool {
	data, err := io.ReadAll(reader)
	if err != nil {
		out.ErrInvalidImageData(fmt.Errorf("%v: reading image data: %w", errCtx, e.ErrValidation))
		return false
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(data))

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

func (i *Interactor) ingestImage(image *im.Image, out OutputPort) bool {
	if err := i.repo.IngestImage(image.Id, image.Collection.Id); err != nil {
		out.ErrInternal(fmt.Errorf("%v: ingesting meta-data: %w", errCtx, e.ErrInternal))
		return false
	}

	for _, label := range image.Labels {
		if err := i.repo.AddLabelToImage(image.Id, image.Collection.Id, label.Label.Id); err != nil {
			out.ErrInternal(fmt.Errorf("%v: ingesting label: %w", errCtx, e.ErrInternal))
			return false
		}
	}

	for _, box := range image.BoundingBoxes {
		if err := i.repo.AddBoundingBoxToImage(image.Id, image.Collection.Id, *box); err != nil {
			out.ErrInternal(fmt.Errorf("%v: ingesting bounding box: %w", errCtx, e.ErrInternal))
			return false
		}
	}
	return true

}

func (i *Interactor) findCollectionByName(name string, out OutputPort) (*clc.Collection, bool) {
	collection, err := i.repo.FindCollectionByName(name)
	baseErrMsg := fmt.Sprintf("%v: finding collection with name %v", errCtx, name)
	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrCollectionNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
		return nil, false
	case errors.Is(err, e.ErrInternal):
		out.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return nil, false
	}
	return collection, true

}

func (i *Interactor) findLabelByName(name string, out OutputPort) (*lbl.Label, bool) {
	baseErrMsg := fmt.Sprintf("%v: fetching label %v", errCtx, name)
	label, err := i.repo.FindLabelByName(name)
	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrLabelNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
		return nil, false
	case errors.Is(err, e.ErrInternal):
		out.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return nil, false
	}
	return label, true

}

func (i *Interactor) duplicateImageExists(hash string, out OutputPort) bool {

	errCtx_ := fmt.Sprintf("%v: searching for duplicate image using hash", errCtx)
	duplicateImage, err := i.repo.FindImageByHash(hash)
	switch {
	case errors.Is(e.ErrNotFound, err):
		return false
	case err == nil:
		out.ErrDuplicateImage(fmt.Errorf("%v: found duplicate with id: %v: %w", errCtx_, duplicateImage.Id, e.ErrValidation))
		return true
	default:
		out.ErrInternal(fmt.Errorf("%v: %w", errCtx_, e.ErrInternal))
		return true
	}
}
