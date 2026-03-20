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
	if ok := i.isUsed(r.Name); !ok {
		return
	}
	if err := i.repo.Delete(r.Name); err != nil {
		i.output.ErrInternal(err)
		return
	}
	i.output.Success()
}

func (i *Interactor) isUsed(name string) bool {
	isUsed, err := i.repo.IsUsed(name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("checking for existence of label with name %v: %w", name, e.ErrInternal))
		return false
	}
	if isUsed {
		i.output.ErrDependency(fmt.Errorf("checking for existence of label with name %v: %w", name, e.ErrDependency))
		return false
	}
	return true

}

func NewInteractor(r Repo, o OutputPort) *Interactor {
	return &Interactor{
		repo:   r,
		output: o,
	}
}
