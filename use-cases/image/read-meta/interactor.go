package read_meta

import (
	"errors"
	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	service st.ImageStore
}

func NewInteractor(service st.ImageStore) *Interactor {
	return &Interactor{service: service}
}

func (i *Interactor) Execute(r Request, out OutputPort) {

	image, err := i.service.Find(im.BaseImage{ImageId: r.ImageId, Collection: r.Collection})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
		default:
			out.ErrInternal(err)
		}
		return
	}

	out.Success(Response{
		Id:         image.Id,
		Collection: image.Collection.Name,
		Labels:     image.Labels})
}
