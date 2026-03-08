package read

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func (i *ReadInteractor) Execute(r ReadRequest) {
	found, err := i.repo.Find(r.Name)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
		default:
			i.output.ErrInternal(err)
		}
		return
	}

	i.output.Success(ReadResponse{Name: found.Name, Description: found.Description})

}

type ReadInteractor struct {
	repo   ReadRepo
	output ReadOutputPort
}

func NewReadInteractor(r ReadRepo, o ReadOutputPort) *ReadInteractor {
	return &ReadInteractor{repo: r, output: o}
}
