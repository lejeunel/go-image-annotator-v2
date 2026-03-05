package delete

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type DeleteInteractor struct {
	repo      DeleteRepo
	presenter DeletePresenter
}

func (i *DeleteInteractor) Execute(r DeleteRequest) {
	if err := i.repo.Delete(DeleteModel{Name: r.Name}); err != nil {
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

func NewDeleteInteractor(r DeleteRepo, p DeletePresenter) *DeleteInteractor {
	return &DeleteInteractor{
		repo:      r,
		presenter: p,
	}
}
