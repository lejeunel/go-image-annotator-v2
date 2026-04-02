package import_shallow

import (
	"errors"

	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Interactor struct {
	repo Repo
}

func NewInteractor(repo Repo) *Interactor {
	return &Interactor{repo: repo}
}
func (i *Interactor) Execute(r Request, out OutputPort) {

	if ok := i.sourceImageExists(r.ImageId, out); !ok {
		return
	}
	collection, ok := i.findDestinationCollection(r.Collection, out)
	if !ok {
		return
	}

	if ok := i.imageDoesNotAlreadyExistInCollection(r.ImageId, collection.Id, out); !ok {
		return
	}

	if err := i.repo.AddImageToCollection(r.ImageId, collection.Id); err != nil {
		out.ErrInternal(err)
		return
	}

	out.Success(Response{})

}
func (i *Interactor) imageDoesNotAlreadyExistInCollection(imageId im.ImageId, collectionId clc.CollectionId, out OutputPort) bool {

	alreadyExists, err := i.repo.ImageExistsInCollection(imageId, collectionId)
	if err != nil {
		out.ErrInternal(err)
		return false
	}
	if alreadyExists {
		out.ErrDependency(e.ErrDependency)
		return false
	}
	return true
}

func (i *Interactor) sourceImageExists(id im.ImageId, out OutputPort) bool {
	exists, err := i.repo.ImageExists(id)
	if err != nil {
		out.ErrInternal(err)
		return false
	}
	if !exists {
		out.ErrNotFound(e.ErrNotFound)
		return false
	}
	return true

}

func (i *Interactor) findDestinationCollection(name string, out OutputPort) (*clc.Collection, bool) {

	collection, err := i.repo.FindCollectionByName(name)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
			return nil, false
		default:
			out.ErrInternal(err)
			return nil, false
		}
	}
	return collection, true

}
