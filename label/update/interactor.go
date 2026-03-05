package update

import (
	"errors"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type UpdateInteractor struct {
	presenter UpdatePresenter
	repo      UpdateRepo
}

func NewUpdateInteractor(r UpdateRepo, p UpdatePresenter) *UpdateInteractor {
	return &UpdateInteractor{repo: r, presenter: p}
}

func (i *UpdateInteractor) Execute(r UpdateRequest) {
	if err := i.repo.Update(UpdateModel{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
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

	i.presenter.Success(UpdateResponse{Name: r.NewName, Description: r.NewDescription})
}
