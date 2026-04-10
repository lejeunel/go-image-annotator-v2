package scroll

import (
	"fmt"
	imstore "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	"log/slog"
)

type Interactor struct {
	store  imstore.Interface
	logger *slog.Logger
	repo   Repo
}

func NewInteractor(store imstore.Interface, repo Repo) *Interactor {
	return &Interactor{store: store, repo: repo, logger: logging.NewNoOpLogger()}
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	var err error
	baseImage, err := i.repo.GetAdjacent(r.ImageId, r.Collection, r.Direction)

	if err != nil {
		i.handleError(err, out)
		return
	}

	image, err := i.store.Find(*baseImage)
	if err != nil {
		i.handleError(err, out)
		return
	}

	out.Success(image)
}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "scrolling images"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)
	out.Error(err)
}
