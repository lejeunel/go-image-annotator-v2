package read

import (
	"fmt"
	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	"log/slog"
)

type Interactor struct {
	store  st.IImageStore
	logger *slog.Logger
}

func NewInteractor(store st.IImageStore) *Interactor {
	return &Interactor{store: store, logger: logging.NewNoOpLogger()}
}

func (i *Interactor) Execute(r Request, out OutputPort) {

	image, err := i.store.Find(im.BaseImage{ImageId: r.ImageId, Collection: r.Collection})
	if err != nil {
		i.handleError(err, out)
		return
	}

	out.Success(image)
}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "reading image meta-data"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)
	out.Error(err)
}
