package remove_bbox

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

func NewInteractor(repo Repo) *Interactor {
	return &Interactor{repo: repo, logger: logging.NewNoOpLogger()}
}
func (i *Interactor) Execute(r Request, out OutputPort) {
	errCtx := "deleting bounding box"
	if err := i.repo.RemoveAnnotation(r.Id); err != nil {
		err = fmt.Errorf("%v: %w", errCtx, err)
		i.logger.Error(errCtx, "error", err)
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
			return
		default:
			out.ErrInternal(err)
			return
		}
	}

	out.Success(Response{})

}
