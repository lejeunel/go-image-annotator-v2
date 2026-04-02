package update

import (
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Interactor struct {
	repo Repo
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{repo: r}
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	exists, err := i.repo.Exists(r.Name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("updating label %v: checking whether it exists: %w", r.Name, e.ErrInternal))
		return
	}
	if !exists {
		out.ErrNotFound(fmt.Errorf("updating label %v: checking whether it exists: %w", r.Name, e.ErrNotFound))
		return
	}

	exists, err = i.repo.Exists(r.NewName)
	if err != nil {
		out.ErrInternal(fmt.Errorf("updating label %v to new name %v: checking whether new name exists: %w", r.Name, r.NewName, e.ErrInternal))
		return
	}
	if exists {
		out.ErrDuplication(fmt.Errorf("updating label %v to new name %v: checking whether new name exists: %w", r.Name, r.NewName, e.ErrDuplicate))
		return
	}
	if err := i.repo.Update(Model{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
		out.ErrInternal(err)
		return
	}

	out.Success(Response{Name: r.NewName, Description: r.NewDescription})
}
