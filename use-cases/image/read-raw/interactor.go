package read_raw

import (
	"errors"
	"io"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output OutputPort
	repo   im.ArtefactReadRepo
}

func (i *Interactor) Execute(r Request) {
	data, err := io.ReadAll(im.NewImageReader(r.ImageId, i.repo))
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
		default:
			i.output.ErrInternal(err)
		}
		return
	}

	i.output.Success(Response{Data: data})

}

func NewInteractor(output OutputPort, repo im.ArtefactReadRepo) *Interactor {
	return &Interactor{output: output, repo: repo}
}
