package read

import (
	"errors"
	c "github.com/lejeunel/go-image-annotator-v2/collection"
)

func (i *ReadCollectionInteractor) Execute(r ReadCollectionRequest) {
	found, err := i.repo.Find(r.Name)
	if err != nil {
		if errors.Is(err, c.ErrNotFound) {
			i.presenter.ErrNotFound(err.Error())
		}
		return
	}

	i.presenter.Success(ReadCollectionResponse{Name: found.Name, Description: found.Description})

}

type ReadCollectionInteractor struct {
	repo      ReadCollectionRepo
	presenter ReadCollectionPresenter
}

func NewReadCollectionInteractor(r ReadCollectionRepo, p ReadCollectionPresenter) *ReadCollectionInteractor {
	return &ReadCollectionInteractor{repo: r, presenter: p}
}
