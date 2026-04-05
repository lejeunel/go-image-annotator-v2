package read_raw

import (
	"errors"
	"fmt"
	"io"

	ast "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	"log/slog"
)

type Interactor struct {
	repo   ast.ArtefactReadRepo
	logger *slog.Logger
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	data, err := io.ReadAll(ast.NewImageReader(r.ImageId, i.repo))
	if err != nil {
		i.handleError(err, out)
		return
	}

	out.Success(Response{Data: data})

}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "reading raw image data"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)

	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(err)
	default:
		out.ErrInternal(err)
	}
}

func NewInteractor(repo ast.ArtefactReadRepo) *Interactor {
	return &Interactor{repo: repo, logger: logging.NewNoOpLogger()}
}
