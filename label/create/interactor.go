package create

import (
	"errors"
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

type Interactor struct {
	output    OutputPort
	repo      Repo
	validator v.Validator
}

func (i *Interactor) checkDuplicate(name string) error {
	errBaseMsg := "checking for duplicate label with name %v: %w"
	alreadyExists, err := i.repo.Exists(name)
	if err != nil {
		return fmt.Errorf(errBaseMsg, name, e.ErrInternal)
	}
	if alreadyExists {
		return fmt.Errorf(errBaseMsg, name, e.ErrDuplicate)
	}
	return nil
}

func (i *Interactor) Execute(r Request) {
	if err := i.validator.Validate(r.Name); err != nil {
		i.output.ErrValidation(err)
		return
	}
	if err := i.checkDuplicate(r.Name); err != nil {
		if errors.Is(err, e.ErrDuplicate) {
			i.output.ErrDuplication(err)
		} else {
			i.output.ErrInternal(err)
		}
		return
	}

	if err := i.repo.Create(Model{Name: r.Name, Description: r.Description}); err != nil {
		i.output.ErrInternal(err)
		return
	}
	i.output.Success(Response{Name: r.Name, Description: r.Description})
}

func NewCreateLabelInteractor(r Repo, v v.Validator, o OutputPort) *Interactor {
	return &Interactor{output: o, repo: r, validator: v}
}
