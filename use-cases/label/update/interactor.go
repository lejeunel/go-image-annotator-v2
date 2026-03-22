package update

import (
	"fmt"

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
	exists, err := i.repo.Exists(r.Name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("updating label %v: checking whether it exists: %w", r.Name, e.ErrInternal))
		return
	}
	if !exists {
		i.output.ErrNotFound(fmt.Errorf("updating label %v: checking whether it exists: %w", r.Name, e.ErrNotFound))
		return
	}

	exists, err = i.repo.Exists(r.NewName)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("updating label %v to new name %v: checking whether new name exists: %w", r.Name, r.NewName, e.ErrInternal))
		return
	}
	if exists {
		i.output.ErrDuplication(fmt.Errorf("updating label %v to new name %v: checking whether new name exists: %w", r.Name, r.NewName, e.ErrDuplicate))
		return
	}
	if err := i.repo.Update(Model{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
		i.output.ErrInternal(err)
		return
	}

	i.output.Success(Response{Name: r.NewName, Description: r.NewDescription})
}
