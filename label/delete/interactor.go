package delete

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type DeleteLabelInteractor struct {
	repo      DeleteLabelRepo
	presenter DeleteLabelPresenter
}

func (i *DeleteLabelInteractor) Execute(r DeleteLabelRequest) {
	if err := i.repo.Delete(DeleteLabelModel{Name: r.Name}); err != nil {
		switch {
		case errors.Is(err, e.ErrDependency):
			i.presenter.ErrDependency(err.Error())
			return
		default:
			i.presenter.ErrInternal(err.Error())
		}
	}
	i.presenter.Success()
}

func NewDeleteLabelInteractor(r DeleteLabelRepo, p DeleteLabelPresenter) *DeleteLabelInteractor {
	return &DeleteLabelInteractor{
		repo:      r,
		presenter: p,
	}
}
