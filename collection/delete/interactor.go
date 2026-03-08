package delete

import (
	"fmt"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type DeleteInteractor struct {
	repo   Repo
	output DeleteOutputPort
}

func (i *DeleteInteractor) Execute(r DeleteRequest) {
	exists, err := i.repo.Exists(r.Name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("deleting collection with name %v: checking whether it exists: %w", r.Name, e.ErrInternal))
		return
	}
	if !exists {
		i.output.ErrNotFound(fmt.Errorf("deleting collection with name %v: checking whether it exists: %w", r.Name, e.ErrNotFound))
		return
	}

	isPopulated, err := i.repo.IsPopulated(r.Name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("deleting collection with name %v: checking whether it contains elements: %w", r.Name, e.ErrInternal))
		return
	}
	if isPopulated {
		i.output.ErrDependency(fmt.Errorf("deleting collection with name %v: checking whether it contains elements: %w", r.Name, e.ErrDependency))
		return
	}

	if err := i.repo.Delete(r.Name); err != nil {
		i.output.ErrInternal(fmt.Errorf("deleting collection with name %v: %w", r.Name, e.ErrInternal))
	}
	i.output.Success()
}

func NewDeleteInteractor(r Repo, o DeleteOutputPort) *DeleteInteractor {
	return &DeleteInteractor{
		repo:   r,
		output: o,
	}
}
