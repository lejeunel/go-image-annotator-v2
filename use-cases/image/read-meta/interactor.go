package read_meta

import (
	"errors"
	"fmt"
	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	"log/slog"
)

type Interactor struct {
	store  st.ImageStore
	logger *slog.Logger
}

func NewInteractor(store st.ImageStore) *Interactor {
	return &Interactor{store: store, logger: logging.NewNoOpLogger()}
}

func (i *Interactor) Execute(r Request, out OutputPort) {

	image, err := i.store.Find(im.BaseImage{ImageId: r.ImageId, Collection: r.Collection})
	if err != nil {
		i.handleError(err, out)
		return
	}

	out.Success(im.Response{
		Id:            image.Id,
		Collection:    image.Collection.Name,
		Labels:        image.Labels,
		BoundingBoxes: image.BoundingBoxes,
	})
}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "reading image meta-data"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)

	switch {
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(err)
	default:
		out.ErrInternal(err)
	}
}
