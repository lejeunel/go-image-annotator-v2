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
	errCtx := "fetching label"
	found, err := i.repo.FindLabelByName(r.Name)

	if err != nil {
		err = fmt.Errorf("%v: %w", errCtx, err)
		i.logger.Error(errCtx, "error", err)

		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
		default:
			out.ErrInternal(err)
		}
		return
	}

	out.Success(Response{Name: found.Name, Description: found.Description})

}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{repo: r, logger: logging.NewNoOpLogger()}
}
