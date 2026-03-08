package create

import (
	"errors"
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

type CreateInteractor struct {
	output    CreateOutputPort
	repo      CreateRepo
	validator v.Validator
}

func (i *CreateInteractor) checkDuplicate(name string) error {
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

func (i *CreateInteractor) Execute(r CreateRequest) {
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

	if err := i.repo.Create(CreateModel{Name: r.Name, Description: r.Description}); err != nil {
		i.output.ErrInternal(err)
		return
	}
	i.output.Success(CreateResponse{Name: r.Name, Description: r.Description})
}

func NewCreateLabelInteractor(r CreateRepo, v v.Validator, o CreateOutputPort) *CreateInteractor {
	return &CreateInteractor{output: o, repo: r, validator: v}
}
