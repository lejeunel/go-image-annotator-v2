package delete

import (
	"errors"
	"fmt"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output  OutputPort
	service im.ImageStore
	repo    Repo
}

func NewInteractor(output OutputPort, service im.ImageStore, repo Repo) *Interactor {
	return &Interactor{output: output, service: service, repo: repo}
}

func (i *Interactor) Execute(r Request) {
	image, ok := i.findImage(r.ImageId, r.Collection)
	if !ok {
		return
	}

	if ok := i.deleteLabels(*image); !ok {
		return
	}

	if ok := i.deleteBoundingBoxes(*image); !ok {
		return
	}

	i.output.Success(Response{})
}
func (i *Interactor) deleteBoundingBoxes(image im.Image) bool {
	errCtx := fmt.Sprintf("deleting bounding boxes assigned to image %v", image.Id.String())
	for _, box := range image.BoundingBoxes {
		if err := i.repo.RemoveAnnotation(image.Id, image.Collection.Id, box.Id); err != nil {
			switch {
			case errors.Is(err, e.ErrNotFound):
				i.output.ErrNotFound(fmt.Errorf("%v: deleting box with label %v: %w", errCtx, box.Label.Name, err))
				return false
			default:
				i.output.ErrInternal(fmt.Errorf("%v: deleting box with label %v: %w", errCtx, box.Label.Name, e.ErrInternal))
				return false
			}
		}
	}
	return true

}

func (i *Interactor) deleteLabels(image im.Image) bool {
	errCtx := fmt.Sprintf("deleting labels assigned to image %v", image.Id.String())
	for _, label := range image.Labels {
		if err := i.repo.RemoveAnnotation(image.Id, image.Collection.Id, label.Id); err != nil {
			switch {
			case errors.Is(err, e.ErrNotFound):
				i.output.ErrNotFound(fmt.Errorf("%v: deleting label %v: %w", errCtx, label.Label.Name, err))
				return false
			default:
				i.output.ErrInternal(fmt.Errorf("%v: deleting label %v: %w", errCtx, label.Label.Name, e.ErrInternal))
				return false
			}
		}
	}
	return true

}

func (i *Interactor) findImage(imageId im.ImageId, collection string) (*im.Image, bool) {
	errCtx := "deleting image: fetching associated resources"
	image, err := i.service.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	switch {
	case errors.Is(err, e.ErrNotFound):
		i.output.ErrNotFound(fmt.Errorf("%v: %w", errCtx, err))
		return nil, false
	case errors.Is(err, e.ErrInternal):
		i.output.ErrInternal(fmt.Errorf("%v: %w", errCtx, e.ErrInternal))
		return nil, false
	}
	return image, true

}
