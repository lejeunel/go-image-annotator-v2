package read

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func (i *ReadInteractor) Execute(r ReadRequest) {
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

	i.presenter.Success(ReadResponse{Name: found.Name, Description: found.Description})

}

type ReadInteractor struct {
	repo      ReadRepo
	presenter ReadPresenter
}

func NewReadInteractor(r ReadRepo, p ReadPresenter) *ReadInteractor {
	return &ReadInteractor{repo: r, presenter: p}
}
