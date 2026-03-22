package create

import (
	"errors"
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

type Interactor struct {
	output    OutputPort
	repo      CreateRepo
	validator v.Validator
}

func (i *Interactor) isDuplicate(name string) bool {
	errBaseMsg := fmt.Sprintf("checking for duplicate collection with name %v", name)
	alreadyExists, err := i.repo.Exists(name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("%v: %w", errBaseMsg, e.ErrInternal))
		return true
	}
	if alreadyExists {
		i.output.ErrDuplication(fmt.Errorf("%v: %w", errBaseMsg, e.ErrDuplicate))
		return true
	}
	return false
}

func (i *Interactor) Execute(r Request) {
	if err := i.validator.Validate(r.Name); err != nil {
		i.output.ErrValidation(err)
		return
	}
	if i.isDuplicate(r.Name) {
		return
	}

	if err := i.repo.Create(Model{Name: r.Name, Description: r.Description}); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			i.output.ErrDuplication(err)
		default:
			i.output.ErrInternal(err)
		}
		return
	}
	i.output.Success(Response{Name: r.Name, Description: r.Description})
}

func NewInteractor(r CreateRepo, v v.Validator, o OutputPort) *Interactor {
	return &Interactor{output: o, repo: r, validator: v}
}
