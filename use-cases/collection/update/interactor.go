package update

import (
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output UpdateOutputPort
	repo   Repo
}

func NewUpdateCollectionInteractor(r Repo, o UpdateOutputPort) *Interactor {
	return &Interactor{repo: r, output: o}
}

func (i *Interactor) Execute(r UpdateRequest) {

	if !i.sourceExists(r.Name) {
		return
	}

	if i.destinationExists(r.NewName) {
		return
	}

	if err := i.repo.Update(UpdateModel{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
		i.output.ErrInternal(fmt.Errorf("updating collection %v: %w", r.Name, e.ErrInternal))
		return
	}

	i.output.Success(UpdateResponse{Name: r.NewName, Description: r.NewDescription})
}

func (i *Interactor) sourceExists(name string) bool {
	baseErrMsg := fmt.Sprintf("updating collection %v: checking whether it exists", name)
	exists, err := i.repo.Exists(name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return true
	}
	if exists {
		return true
	}
	i.output.ErrNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
	return false
}

func (i *Interactor) destinationExists(name string) bool {
	baseErrMsg := fmt.Sprintf("updating collection to name %v: checking whether it exists", name)
	exists, err := i.repo.Exists(name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return true
	}
	if exists {
		i.output.ErrDuplication(fmt.Errorf("%v: %w", baseErrMsg, e.ErrDuplicate))
		return true
	}
	return false
}
