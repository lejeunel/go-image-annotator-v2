package delete

import (
	"fmt"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo Repo
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	errCtx := fmt.Sprintf("deleting collection with name %v", r.Name)
	if ok := i.exists(r, out, errCtx); !ok {
		return
	}
	if ok := i.isDeletable(r, out, errCtx); !ok {
		return
	}

	if err := i.repo.Delete(r.Name); err != nil {
		out.ErrInternal(fmt.Errorf("%v: %w", errCtx, e.ErrInternal))
		return
	}
	out.Success()
}
func (i *Interactor) isDeletable(r Request, out OutputPort, errCtx string) bool {
	isPopulated, err := i.repo.IsPopulated(r.Name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("%v: checking whether it contains items: %w", errCtx, e.ErrInternal))
		return false
	}
	if *isPopulated {
		out.ErrDependency(fmt.Errorf("%v: checking whether it contains items: %w", errCtx, e.ErrDependency))
		return false
	}
	return true

}

func (i *Interactor) exists(r Request, out OutputPort, errCtx string) bool {
	exists, err := i.repo.Exists(r.Name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("%v: checking whether it exists: %w", errCtx, e.ErrInternal))
		return false
	}
	if !exists {
		out.ErrNotFound(fmt.Errorf("%v: checking whether it exists: %w", errCtx, e.ErrNotFound))
		return false
	}
	return true
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{
		repo: r,
	}
}
