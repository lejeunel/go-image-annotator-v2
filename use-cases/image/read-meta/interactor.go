package read_meta

import (
	"errors"
	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Interactor struct {
	store st.ImageStore
}

func NewInteractor(store st.ImageStore) *Interactor {
	return &Interactor{store: store}
}

func (i *Interactor) Execute(r Request, out OutputPort) {

	image, err := i.store.Find(im.BaseImage{ImageId: r.ImageId, Collection: r.Collection})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
		default:
			out.ErrInternal(err)
		}
		return
	}

	out.Success(im.Response{
		Id:            image.Id,
		Collection:    image.Collection.Name,
		Labels:        image.Labels,
		BoundingBoxes: image.BoundingBoxes,
	})
}
