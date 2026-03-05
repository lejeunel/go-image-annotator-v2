package update

import (
	"errors"
	c "github.com/lejeunel/go-image-annotator-v2/collection"
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
		if errors.Is(err, c.ErrDuplicate) {
			i.presenter.ErrDuplication(err.Error())
		}
		return
	}

	i.presenter.Success(UpdateCollectionResponse{Name: r.Name, Description: r.Description})
}
