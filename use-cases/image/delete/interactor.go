package delete

import (
	"errors"
	"fmt"

	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	service st.ImageStore
	repo    Repo
}

func NewInteractor(service st.ImageStore, repo Repo) *Interactor {
	return &Interactor{service: service, repo: repo}
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	image, ok := i.findImage(r.ImageId, r.Collection, out)
	if !ok {
		return
	}

	if ok := i.deleteLabels(*image, out); !ok {
		return
	}

	if ok := i.deleteBoundingBoxes(*image, out); !ok {
		return
	}

	if err := i.repo.RemoveImageFromCollection(image.Id, image.Collection.Id); err != nil {
		out.ErrInternal(fmt.Errorf("removing image %v from collection %v: %w",
			image.Id, image.Collection.Name, e.ErrInternal))
		return

	}

	out.Success(Response{})
}
func (i *Interactor) deleteBoundingBoxes(image im.Image, out OutputPort) bool {
	errCtx := fmt.Sprintf("deleting bounding boxes assigned to image %v", image.Id.String())
	for _, box := range image.BoundingBoxes {
		if err := i.repo.RemoveAnnotation(image.Id, image.Collection.Id, box.Id); err != nil {
			switch {
			case errors.Is(err, e.ErrNotFound):
				out.ErrNotFound(fmt.Errorf("%v: deleting box with label %v: %w", errCtx, box.Label.Name, err))
				return false
			default:
				out.ErrInternal(fmt.Errorf("%v: deleting box with label %v: %w", errCtx, box.Label.Name, e.ErrInternal))
				return false
			}
		}
	}
	return true

}

func (i *Interactor) deleteLabels(image im.Image, out OutputPort) bool {
	errCtx := fmt.Sprintf("deleting labels assigned to image %v", image.Id.String())
	for _, label := range image.Labels {
		if err := i.repo.RemoveAnnotation(image.Id, image.Collection.Id, label.Id); err != nil {
			switch {
			case errors.Is(err, e.ErrNotFound):
				out.ErrNotFound(fmt.Errorf("%v: deleting label %v: %w", errCtx, label.Label.Name, err))
				return false
			default:
				out.ErrInternal(fmt.Errorf("%v: deleting label %v: %w", errCtx, label.Label.Name, e.ErrInternal))
				return false
			}
		}
	}
	return true

}

func (i *Interactor) findImage(imageId im.ImageId, collection string, out OutputPort) (*im.Image, bool) {
	errCtx := "deleting image: fetching associated resources"
	image, err := i.service.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(fmt.Errorf("%v: %w", errCtx, err))
		return nil, false
	case errors.Is(err, e.ErrInternal):
		out.ErrInternal(fmt.Errorf("%v: %w", errCtx, e.ErrInternal))
		return nil, false
	}
	return image, true

}
