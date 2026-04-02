package remove_bbox

import (
	"errors"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Interactor struct {
	repo Repo
}

func NewInteractor(repo Repo) *Interactor {
	return &Interactor{repo: repo}
}
func (i *Interactor) Execute(r Request, out OutputPort) {
	if err := i.repo.RemoveAnnotation(r.Id); err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(e.ErrNotFound)
			return
		default:
			out.ErrInternal(e.ErrInternal)
			return
		}
	}

	out.Success(Response{})

}
