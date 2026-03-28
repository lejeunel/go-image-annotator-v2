package read_raw

import (
	"errors"
	"io"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo im.ArtefactReadRepo
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	data, err := io.ReadAll(im.NewImageReader(r.ImageId, i.repo))
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
		default:
			out.ErrInternal(err)
		}
		return
	}

	out.Success(Response{Data: data})

}

func NewInteractor(repo im.ArtefactReadRepo) *Interactor {
	return &Interactor{repo: repo}
}
