package delete

import (
	"errors"
	"fmt"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	"log/slog"
)

type Interactor struct {
	repo   Repo
	logger *slog.Logger
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	if err := i.isUsed(r.Name); err != nil {
		i.handleError(err, out)
		return
	}
	if err := i.exists(r.Name); err != nil {
		i.handleError(err, out)
		return
	}

	if err := i.repo.Delete(r.Name); err != nil {
		i.handleError(err, out)
		return
	}
	out.Success()
}
func (i *Interactor) exists(name string) error {
	baseErr := fmt.Errorf("checking whether label with name %v exists", name)
	exists, err := i.repo.Exists(name)
	if err != nil {
		return fmt.Errorf("%w: %v: %w", baseErr, err, e.ErrInternal)
	}
	if !exists {
		return fmt.Errorf("%w: %v: %w", baseErr, err, e.ErrNotFound)
	}
	return nil
}

func (i *Interactor) isUsed(name string) error {
	baseErr := fmt.Errorf("checking whether label with name %v is used", name)
	isUsed, err := i.repo.IsUsed(name)
	if err != nil {
		return fmt.Errorf("%w: %w", baseErr, e.ErrInternal)
	}
	if *isUsed {
		return fmt.Errorf("%w: %w", baseErr, e.ErrDependency)
	}
	return nil

}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "creating label"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)

	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(err)
	case errors.Is(err, e.ErrDependency):
		out.ErrDependency(err)
	default:
		out.ErrInternal(err)
	}
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{repo: r, logger: logging.NewNoOpLogger()}
}
