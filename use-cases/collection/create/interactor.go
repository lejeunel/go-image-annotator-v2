package create

import (
	"errors"
	"fmt"
	"log/slog"

	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	v "github.com/lejeunel/go-image-annotator-v2/shared/validation"
)

type Interactor struct {
	repo      CreateRepo
	validator v.Validator
	logger    *slog.Logger
}

func (i *Interactor) Execute(r Request, out OutputPort) {

	if err := i.validate(r.Name); err != nil {
		i.handleError(err, out)
		return
	}

	if err := i.create(r); err != nil {
		i.handleError(err, out)
		return
	}

	out.Success(Response{Name: r.Name, Description: r.Description})
}

func (i *Interactor) create(r Request) error {
	collection := clc.NewCollection(clc.NewCollectionId(), r.Name, clc.WithDescription(r.Description))
	if err := i.repo.Create(*collection); err != nil {
		return err
	}
	return nil

}

func (i *Interactor) validate(name string) error {
	if err := i.validator.Validate(name); err != nil {
		return fmt.Errorf("checking collection name %v: %w", name, err)
	}
	if err := i.isDuplicate(name); err != nil {
		return err
	}
	return nil

}

func (i *Interactor) isDuplicate(name string) error {
	errBaseMsg := fmt.Sprintf("checking for duplicate collection with name %v", name)
	alreadyExists, err := i.repo.Exists(name)
	if err != nil {
		return fmt.Errorf("%v: %w", errBaseMsg, e.ErrInternal)
	}
	if alreadyExists {
		return fmt.Errorf("%v: %w", errBaseMsg, e.ErrDuplicate)
	}
	return nil
}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "creating collection"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)

	switch {
	case errors.Is(err, e.ErrValidation):
		out.ErrValidation(err)
	case errors.Is(err, e.ErrDuplicate):
		out.ErrDuplication(err)
	default:
		out.ErrInternal(err)
	}
}

func NewInteractor(r CreateRepo, v v.Validator) *Interactor {
	return &Interactor{repo: r, validator: v, logger: logging.NewNoOpLogger()}
}
