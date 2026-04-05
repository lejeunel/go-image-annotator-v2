package unassign_label

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
	repo   Repo
	output OutputPort
	store  st.ImageStore
	logger *slog.Logger
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	image, err := i.findImage(r.ImageId, r.Collection)
	if err != nil {
		i.handleError(err, out)
		return
	}

	if err := i.removeLabel(*image, r.Label); err != nil {
		i.handleError(err, out)
		return
	}

	out.Success(Response{})

}

func (i *Interactor) findImage(imageId im.ImageId, collection string) (*im.Image, error) {
	baseErr := fmt.Errorf("finding image %v in collection %v", imageId, collection)
	image, err := i.store.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", baseErr, err)
	}
	return image, nil
}

func (i *Interactor) removeLabel(image im.Image, label string) error {
	baseErr := fmt.Errorf("removing image label")
	removed := 0
	for _, imageLabel := range image.Labels {
		if imageLabel.Label.Name == label {
			err := i.repo.RemoveImageLabel(image.Id, image.Collection.Id, imageLabel.Label.Id)
			if err != nil {
				return fmt.Errorf("%w: %w", baseErr, err)
			}
			removed += 1
		}
	}

	if removed == 0 {
		return fmt.Errorf("%w: %w", baseErr, e.ErrNotFound)
	}
	return nil

}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "deleting image label"
	err = fmt.Errorf("%v: %w", errCtx, err)
	i.logger.Error(errCtx, "error", err)

	switch {
	case errors.Is(err, e.ErrDependency):
		out.ErrDependency(err)
	case errors.Is(err, e.ErrNotFound):
		out.ErrNotFound(err)
	default:
		out.ErrInternal(err)
	}
}

func NewInteractor(repo Repo, store st.ImageStore) *Interactor {
	return &Interactor{repo: repo, store: store, logger: logging.NewNoOpLogger()}
}
