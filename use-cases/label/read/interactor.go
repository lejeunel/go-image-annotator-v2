package read

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func (i *Interactor) Execute(r Request) {
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

	i.output.Success(Response{Name: found.Name, Description: found.Description})

}

type Interactor struct {
	repo   Repo
	output OutputPort
}

func NewReadInteractor(r Repo, o OutputPort) *Interactor {
	return &Interactor{repo: r, output: o}
}
