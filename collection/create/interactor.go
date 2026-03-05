package create

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type CreateCollectionInteractor struct {
	presenter CreateCollectionPresenter
	repo      CreateCollectionRepo
}

func (i *CreateCollectionInteractor) Execute(r CreateCollectionRequest) {
	if err := i.repo.Create(r); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			i.presenter.ErrDuplication(err.Error())
		default:
			i.presenter.ErrInternal(err.Error())
		}
	}
	i.presenter.Success(CreateCollectionResponse{Name: r.Name, Description: r.Description})
}

func NewCreateCollectionInteractor(r CreateCollectionRepo, p CreateCollectionPresenter) *CreateCollectionInteractor {
	return &CreateCollectionInteractor{presenter: p, repo: r}
}
