package create

import (
	"errors"
	"fmt"

	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	v "github.com/lejeunel/go-image-annotator-v2/validation"
)

type Interactor struct {
	repo      Repo
	validator v.Validator
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	if err := i.validator.Validate(r.Name); err != nil {
		out.ErrValidation(err)
		return
	}
	if err := i.checkDuplicate(r.Name, out); err != nil {
		if errors.Is(err, e.ErrDuplicate) {
			out.ErrDuplication(err)
		} else {
			out.ErrInternal(err)
		}
		return
	}

	label := lbl.NewLabel(lbl.NewLabelId(), r.Name, lbl.WithDescription(r.Description))
	if err := i.repo.Create(*label); err != nil {
		out.ErrInternal(err)
		return
	}
	out.Success(Response{Name: r.Name, Description: r.Description})
}

func (i *Interactor) checkDuplicate(name string, out OutputPort) error {
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

func NewInteractor(r Repo, v v.Validator) *Interactor {
	return &Interactor{repo: r, validator: v}
}
