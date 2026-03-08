package update

import (
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type UpdateInteractor struct {
	output UpdateOutputPort
	repo   Repo
}

func NewUpdateCollectionInteractor(r Repo, o UpdateOutputPort) *UpdateInteractor {
	return &UpdateInteractor{repo: r, output: o}
}

func (i *UpdateInteractor) Execute(r UpdateRequest) {
	exists, err := i.repo.Exists(r.Name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("updating collection %v: checking whether it exists: %w", r.Name, e.ErrInternal))
		return
	}
	if !exists {
		i.output.ErrNotFound(fmt.Errorf("updating collection %v: checking whether it exists: %w", r.Name, e.ErrNotFound))
		return
	}

	exists, err = i.repo.Exists(r.NewName)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("updating collection %v to new name %v: checking whether new name exists: %w", r.Name, r.NewName, e.ErrInternal))
		return
	}
	if exists {
		i.output.ErrDuplication(fmt.Errorf("updating collection %v to new name %v: checking whether new name exists: %w", r.Name, r.NewName, e.ErrDuplicate))
		return
	}

	if err := i.repo.Update(UpdateModel{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
		i.output.ErrInternal(fmt.Errorf("updating collection %v: %w", r.Name, err))
		return
	}

	i.output.Success(UpdateResponse{Name: r.NewName, Description: r.NewDescription})
}
