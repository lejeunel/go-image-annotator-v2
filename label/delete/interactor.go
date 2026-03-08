package delete

import (
	"errors"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type DeleteInteractor struct {
	repo   DeleteRepo
	output DeleteOutputPort
}

func (i *DeleteInteractor) Execute(r DeleteRequest) {
	if err := i.repo.Delete(DeleteModel{Name: r.Name}); err != nil {
		switch {
		case errors.Is(err, e.ErrDependency):
			i.output.ErrDependency(err)
			return
		default:
			i.output.ErrInternal(err)
		}
	}
	i.output.Success()
}

func NewDeleteLabelInteractor(r DeleteRepo, o DeleteOutputPort) *DeleteInteractor {
	return &DeleteInteractor{
		repo:   r,
		output: o,
	}
}
