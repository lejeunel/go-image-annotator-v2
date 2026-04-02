package read

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func (i *Interactor) Execute(r Request, out OutputPort) {
	found, err := i.repo.FindCollectionByName(r.Name)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
		default:
			out.ErrInternal(err)
		}
		return
	}

	out.Success(Response{Name: found.Name, Description: found.Description})

}

type Interactor struct {
	repo Repo
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{repo: r}
}
