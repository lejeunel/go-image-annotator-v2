package delete

import (
	"fmt"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo Repo
}

func (i *Interactor) Execute(r DeleteRequest, out OutputPort) {
	exists, err := i.repo.Exists(r.Name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("deleting collection with name %v: checking whether it exists: %w", r.Name, e.ErrInternal))
		return
	}
	if !exists {
		out.ErrNotFound(fmt.Errorf("deleting collection with name %v: checking whether it exists: %w", r.Name, e.ErrNotFound))
		return
	}

	isPopulated, err := i.repo.IsPopulated(r.Name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("deleting collection with name %v: checking whether it contains elements: %w", r.Name, e.ErrInternal))
		return
	}
	if isPopulated {
		out.ErrDependency(fmt.Errorf("deleting collection with name %v: checking whether it contains elements: %w", r.Name, e.ErrDependency))
		return
	}

	if err := i.repo.Delete(r.Name); err != nil {
		out.ErrInternal(fmt.Errorf("deleting collection with name %v: %w", r.Name, e.ErrInternal))
	}
	out.Success()
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{
		repo: r,
	}
}
