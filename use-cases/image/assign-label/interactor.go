package assign_label

import (
	"errors"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo   Repo
	output OutputPort
	store  im.ImageStore
}

func (i *Interactor) Execute(r Request) {
	_, err := i.store.Find(im.BaseImage{ImageId: r.ImageId, Collection: r.Collection})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
		case errors.Is(err, e.ErrDependency):
			i.output.ErrDependency(err)
		default:
			i.output.ErrInternal(err)
		}
		return
	}

	i.output.Success(Response{ImageId: r.ImageId, Collection: r.Collection, Label: r.Label})
}

func (i *Interactor) addLabel(imageId im.ImageId, collectionId clc.CollectionId, labelId lbl.LabelId) bool {
	err := i.repo.AddLabel(imageId, collectionId, labelId)
	if err != nil {
		i.output.ErrInternal(err)
		return false
	}
	return true

}

func NewInteractor(repo Repo, output OutputPort, store im.ImageStore) *Interactor {
	return &Interactor{repo: repo, output: output, store: store}
}
