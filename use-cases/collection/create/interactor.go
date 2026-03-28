package create

import (
	"errors"
	"fmt"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

type Interactor struct {
	repo      CreateRepo
	validator v.Validator
}

func (i *Interactor) Execute(r Request, out OutputPort) {

	if ok := i.validate(r.Name, out); !ok {
		return
	}

	if ok := i.create(r, out); !ok {
		return
	}

	out.Success(Response{Name: r.Name, Description: r.Description})
}

func (i *Interactor) create(r Request, out OutputPort) bool {
	collection := clc.NewCollection(clc.NewCollectionId(), r.Name, clc.WithDescription(r.Description))
	if err := i.repo.Create(*collection); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			out.ErrDuplication(err)
			return false
		default:
			out.ErrInternal(err)
			return false
		}
	}
	return true

}

func (i *Interactor) validate(name string, out OutputPort) bool {
	if err := i.validator.Validate(name); err != nil {
		out.ErrValidation(err)
		return false
	}
	if i.isDuplicate(name, out) {
		return false
	}
	return true

}

func (i *Interactor) isDuplicate(name string, out OutputPort) bool {
	errBaseMsg := fmt.Sprintf("checking for duplicate collection with name %v", name)
	alreadyExists, err := i.repo.Exists(name)
	if err != nil {
		out.ErrInternal(fmt.Errorf("%v: %w", errBaseMsg, e.ErrInternal))
		return true
	}
	if alreadyExists {
		out.ErrDuplication(fmt.Errorf("%v: %w", errBaseMsg, e.ErrDuplicate))
		return true
	}
	return false
}

func NewInteractor(r CreateRepo, v v.Validator) *Interactor {
	return &Interactor{repo: r, validator: v}
}
