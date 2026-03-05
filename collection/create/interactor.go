package create

import (
	"errors"
	c "github.com/lejeunel/go-image-annotator-v2/collection"
)

type CreateCollectionInteractor struct {
	presenter CreateCollectionPresenter
	repo      CreateCollectionRepo
}

func (i *CreateCollectionInteractor) Execute(r CreateCollectionRequest) {
	if err := i.repo.Create(r); err != nil {
		if errors.Is(err, c.ErrDuplicate) {
			i.presenter.ErrDuplication(err.Error())
		}
		return
	}
	i.presenter.Success(CreateCollectionResponse{Name: r.Name, Description: r.Description})
}

func NewCreateCollectionInteractor(r CreateCollectionRepo, p CreateCollectionPresenter) *CreateCollectionInteractor {
	return &CreateCollectionInteractor{presenter: p, repo: r}
}
