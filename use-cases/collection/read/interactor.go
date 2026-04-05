package read

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
	found, err := i.repo.FindCollectionByName(r.Name)
	if err != nil {
		i.handleError(err, out)
		return
	}

	out.Success(Response{Name: found.Name, Description: found.Description})

}
func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "fetching collection"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)

	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(err)
	default:
		out.ErrInternal(err)
	}
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{
		repo:   r,
		logger: logging.NewNoOpLogger(),
	}
}
