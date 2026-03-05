package create

import (
	"errors"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type CreateLabelInteractor struct {
	presenter CreateLabelPresenter
	repo      CreateRepo
}

func (i *CreateLabelInteractor) Execute(r CreateLabelRequest) {
	if err := i.repo.Create(CreateModel{Name: r.Name, Description: r.Description}); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			i.presenter.ErrDuplication(err.Error())
		default:
			i.presenter.ErrInternal(err.Error())
		}
	}
	i.presenter.Success(CreateLabelResponse{Name: r.Name, Description: r.Description})
}

func NewCreateLabelInteractor(r CreateRepo, p CreateLabelPresenter) *CreateLabelInteractor {
	return &CreateLabelInteractor{presenter: p, repo: r}
}
