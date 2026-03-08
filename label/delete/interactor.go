package delete

import (
	"fmt"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo   Repo
	output OutputPort
}

func (i *Interactor) Execute(r Request) {
	isUsed, err := i.repo.IsUsed(r.Name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("checking for existence of label with name %v: %w", r.Name, e.ErrInternal))
		return
	}
	if isUsed {
		i.output.ErrDependency(fmt.Errorf("checking for existence of label with name %v: %w", r.Name, e.ErrDependency))
		return
	}
	if err := i.repo.Delete(r.Name); err != nil {
		i.output.ErrInternal(err)
		return
	}
	i.output.Success()
}

func NewDeleteLabelInteractor(r Repo, o OutputPort) *Interactor {
	return &Interactor{
		repo:   r,
		output: o,
	}
}
