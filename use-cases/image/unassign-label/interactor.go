package unassign_label

import (
	"errors"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo   Repo
	output OutputPort
	store  im.ImageStore
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	image, ok := i.findImage(r.ImageId, r.Collection, out)
	if !ok {
		return
	}

	if ok := i.removeLabel(*image, r.Label, out); !ok {
		return
	}

	out.Success(Response{})

}

func (i *Interactor) findImage(imageId im.ImageId, collection string, out OutputPort) (*im.Image, bool) {
	image, err := i.store.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
			return nil, false
		case errors.Is(err, e.ErrDependency):
			out.ErrDependency(err)
			return nil, false
		default:
			out.ErrInternal(err)
			return nil, false
		}
	}
	return image, true
}

func (i *Interactor) removeLabel(image im.Image, label string, out OutputPort) bool {
	removed := 0
	for _, imageLabel := range image.Labels {
		if imageLabel.Label.Name == label {
			err := i.repo.RemoveLabel(image.Id, image.Collection.Id, imageLabel.Label.Id)
			if err != nil {
				out.ErrInternal(err)
				return false
			}
			removed += 1
		}
	}

	if removed == 0 {
		out.ErrNotFound(e.ErrNotFound)
		return false
	}
	return true

}

func NewInteractor(repo Repo, store im.ImageStore) *Interactor {
	return &Interactor{repo: repo, store: store}
}
