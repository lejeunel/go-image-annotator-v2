package read

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func (i *ReadCollectionInteractor) Execute(r ReadCollectionRequest) {
	found, err := i.repo.Find(r.Name)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.presenter.ErrNotFound(err.Error())
		default:
			i.presenter.ErrInternal(err.Error())
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
