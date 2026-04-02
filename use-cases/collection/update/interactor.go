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

	if !i.sourceExists(r.Name, out) {
		return
	}

	if i.destinationExists(r.NewName, out) {
		return
	}

	if err := i.repo.Update(Model{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
		out.ErrInternal(fmt.Errorf("updating collection %v: %w", r.Name, e.ErrInternal))
		return
	}

	out.Success(Response{Name: r.NewName, Description: r.NewDescription})
}

func (i *Interactor) sourceExists(name string, out OutputPort) bool {
	baseErrMsg := fmt.Sprintf("updating collection %v: checking whether it exists", name)
	exists, err := i.repo.Exists(name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return true
	}
	if exists {
		return true
	}
	out.ErrNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
	return false
}

func (i *Interactor) destinationExists(name string, out OutputPort) bool {
	baseErrMsg := fmt.Sprintf("updating collection to name %v: checking whether it exists", name)
	exists, err := i.repo.Exists(name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return true
	}
	if exists {
		out.ErrDuplication(fmt.Errorf("%v: %w", baseErrMsg, e.ErrDuplicate))
		return true
	}
	return false
}
