package delete

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo   Repo
	output OutputPort
}

func (i *Interactor) Execute(r Request) {
	if err := i.repo.Delete(Model{Name: r.Name}); err != nil {
		switch {
		case errors.Is(err, e.ErrDependency):
			i.output.ErrDependency(err)
			return
		default:
			i.output.ErrInternal(err)
		}
	}
	i.output.Success()
}

func NewDeleteLabelInteractor(r Repo, o OutputPort) *Interactor {
	return &Interactor{
		repo:   r,
		output: o,
	}
}
