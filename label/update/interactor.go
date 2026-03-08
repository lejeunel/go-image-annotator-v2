package update

import (
	"errors"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type UpdateInteractor struct {
	output UpdateOutputPort
	repo   UpdateRepo
}

func NewUpdateInteractor(r UpdateRepo, o UpdateOutputPort) *UpdateInteractor {
	return &UpdateInteractor{repo: r, output: o}
}

func (i *UpdateInteractor) Execute(r UpdateRequest) {
	if err := i.repo.Update(UpdateModel{Name: r.Name, NewName: r.NewName, NewDescription: r.NewDescription}); err != nil {
		switch {
		case errors.Is(err, e.ErrDuplicate):
			i.output.ErrDuplication(err.Error())
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err.Error())
		default:
			i.output.ErrInternal(err.Error())
		}
		return
	}

	i.output.Success(UpdateResponse{Name: r.NewName, Description: r.NewDescription})
}
