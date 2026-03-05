package update

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type UpdateCollectionInteractor struct {
	presenter UpdateCollectionPresenter
	repo      UpdateCollectionRepo
}

func NewUpdateCollectionInteractor(r UpdateCollectionRepo, p UpdateCollectionPresenter) *UpdateCollectionInteractor {
	return &UpdateCollectionInteractor{repo: r, presenter: p}
}

func (i *UpdateCollectionInteractor) Execute(r UpdateCollectionRequest) {
	if err := i.repo.Update(r); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			i.presenter.ErrDuplication(err.Error())
		case errors.Is(err, e.ErrNotFound):
			i.presenter.ErrNotFound(err.Error())
		default:
			i.presenter.ErrInternal(err.Error())
		}
		return
	}

	i.presenter.Success(UpdateCollectionResponse{Name: r.NewName, Description: r.NewDescription})
}
