package delete

import (
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Interactor struct {
	repo Repo
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	if ok := i.isUsed(r.Name, out); !ok {
		return
	}
	if ok := i.exists(r.Name, out); !ok {
		return
	}

	if err := i.repo.Delete(r.Name); err != nil {
		out.ErrInternal(err)
		return
	}
	out.Success()
}
func (i *Interactor) exists(name string, out OutputPort) bool {
	exists, err := i.repo.Exists(name)
	if err != nil {
		out.ErrInternal(err)
		return false
	}
	if !exists {
		out.ErrNotFound(e.ErrNotFound)
		return false
	}
	return true
}

func (i *Interactor) isUsed(name string, out OutputPort) bool {
	isUsed, err := i.repo.IsUsed(name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("checking for existence of label with name %v: %w", name, e.ErrInternal))
		return false
	}
	if *isUsed {
		out.ErrDependency(fmt.Errorf("checking for existence of label with name %v: %w", name, e.ErrDependency))
		return false
	}
	return true

}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{repo: r}
}
