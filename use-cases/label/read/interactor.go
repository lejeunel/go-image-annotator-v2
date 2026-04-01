package read

import (
	"errors"
	"fmt"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

func (i *Interactor) Execute(r Request, out OutputPort) {
	errCtx := "finding label by name"
	found, err := i.repo.FindLabelByName(r.Name)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(fmt.Errorf("%v: %w", errCtx, err))
		default:
			out.ErrInternal(fmt.Errorf("%v: %w", errCtx, err))
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
