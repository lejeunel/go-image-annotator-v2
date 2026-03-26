package create

import (
	"errors"
	"fmt"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

type Interactor struct {
	output    OutputPort
	repo      CreateRepo
	validator v.Validator
}

func (i *Interactor) Execute(r Request) {

	if ok := i.validate(r.Name); !ok {
		return
	}

	if ok := i.create(r); !ok {
		return
	}

	i.output.Success(Response{Name: r.Name, Description: r.Description})
}

func (i *Interactor) create(r Request) bool {
	collection := clc.NewCollection(r.Name, clc.WithDescription(r.Description))
	if err := i.repo.Create(*collection); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			i.output.ErrDuplication(err)
			return false
		default:
			i.output.ErrInternal(err)
			return false
		}
	}
	return true

}

func (i *Interactor) validate(name string) bool {
	if err := i.validator.Validate(name); err != nil {
		i.output.ErrValidation(err)
		return false
	}
	if i.isDuplicate(name) {
		return false
	}
	return true

}

func (i *Interactor) isDuplicate(name string) bool {
	errBaseMsg := fmt.Sprintf("checking for duplicate collection with name %v", name)
	alreadyExists, err := i.repo.CollectionWithNameExists(name)
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

func NewInteractor(r CreateRepo, v v.Validator, o OutputPort) *Interactor {
	return &Interactor{output: o, repo: r, validator: v}
}
