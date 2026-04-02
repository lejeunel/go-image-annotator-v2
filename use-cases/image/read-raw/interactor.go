package read_raw

import (
	"errors"
	"io"

	ast "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Interactor struct {
	repo ast.ArtefactReadRepo
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	data, err := io.ReadAll(ast.NewImageReader(r.ImageId, i.repo))
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

func NewInteractor(repo ast.ArtefactReadRepo) *Interactor {
	return &Interactor{repo: repo}
}
