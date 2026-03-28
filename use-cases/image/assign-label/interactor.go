package assign_label

import (
	"errors"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo  Repo
	store im.ImageStore
}

func (i *Interactor) Execute(r Request, out OutputPort) {

	image, ok := i.findImage(r.ImageId, r.Collection, out)
	if !ok {
		return
	}
	label, ok := i.findLabel(r.Label, out)
	if !ok {
		return
	}

	if ok := i.addLabel(image.Id, image.Collection.Id, label.Id, out); !ok {
		return
	}

	out.Success(Response{ImageId: r.ImageId, Collection: r.Collection, Label: r.Label})
}
func (i *Interactor) findLabel(name string, out OutputPort) (*lbl.Label, bool) {
	label, err := i.repo.FindLabel(name)
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
	return label, true

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

func (i *Interactor) addLabel(imageId im.ImageId, collectionId clc.CollectionId, labelId lbl.LabelId, out OutputPort) bool {
	err := i.repo.AddLabel(imageId, collectionId, labelId)
	if err != nil {
		out.ErrInternal(err)
		return false
	}
	return true

}

func NewInteractor(repo Repo, store im.ImageStore) *Interactor {
	return &Interactor{repo: repo, store: store}
}
