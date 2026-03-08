package update

import (
	"errors"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output OutputPort
	repo   Repo
}

func NewUpdateInteractor(r Repo, o OutputPort) *Interactor {
	return &Interactor{repo: r, output: o}
}

func (i *Interactor) Execute(r Request) {
	if err := i.repo.Update(Model{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			i.output.ErrDuplication(err.Error())
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err.Error())
		default:
			i.output.ErrInternal(err.Error())
		}
		return
	}

	i.output.Success(Response{Name: r.NewName, Description: r.NewDescription})
}
