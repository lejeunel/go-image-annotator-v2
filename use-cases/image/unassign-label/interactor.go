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

func (i *Interactor) Execute(r Request) {
	image, ok := i.findImage(r.ImageId, r.Collection)
	if !ok {
		return
	}

	if ok := i.removeLabel(*image, r.Label); !ok {
		return
	}

	i.output.Success(Response{})

}

func (i *Interactor) findImage(imageId im.ImageId, collection string) (*im.Image, bool) {
	image, err := i.store.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
			return nil, false
		case errors.Is(err, e.ErrDependency):
			i.output.ErrDependency(err)
			return nil, false
		default:
			i.output.ErrInternal(err)
			return nil, false
		}
	}
	return image, true
}

func (i *Interactor) removeLabel(image im.Image, label string) bool {
	removed := 0
	for _, imageLabel := range image.Labels {
		if imageLabel.Label.Name == label {
			err := i.repo.RemoveLabel(image.Id, image.Collection.Id, imageLabel.Label.Id)
			if err != nil {
				i.output.ErrInternal(err)
				return false
			}
			removed += 1
		}
	}

	if removed == 0 {
		i.output.ErrNotFound(e.ErrNotFound)
		return false
	}
	return true

}

func NewInteractor(repo Repo, output OutputPort, store im.ImageStore) *Interactor {
	return &Interactor{repo: repo, output: output, store: store}
}
