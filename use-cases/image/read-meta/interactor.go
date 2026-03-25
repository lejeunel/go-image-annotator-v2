package read_meta

import (
	"errors"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output  OutputPort
	service im.ImageStore
}

func NewInteractor(output OutputPort, service im.ImageStore) *Interactor {
	return &Interactor{output: output, service: service}
}

func (i *Interactor) Execute(r Request) {

	image, err := i.service.Find(im.BaseImage{ImageId: r.ImageId, Collection: r.Collection})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
		default:
			i.output.ErrInternal(err)
		}
		return
	}

	i.output.Success(Response{
		Id:         image.Id,
		Collection: image.Collection.Name,
		Labels:     image.Labels})
}
