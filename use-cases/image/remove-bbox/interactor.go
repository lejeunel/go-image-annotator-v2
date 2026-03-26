package remove_bbox

import (
	"errors"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output OutputPort
	repo   Repo
}

func NewInteractor(output OutputPort, repo Repo) *Interactor {
	return &Interactor{output: output, repo: repo}
}
func (i *Interactor) Execute(r Request) {
	if err := i.repo.RemoveAnnotation(r.Id); err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(e.ErrNotFound)
			return
		default:
			i.output.ErrInternal(e.ErrInternal)
			return
		}
	}

	i.output.Success(Response{})

}
