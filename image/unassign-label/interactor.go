package unassign_label

import (
	"errors"
	labeling "github.com/lejeunel/go-image-annotator-v2/domain/labeling"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo            Repo
	output          OutputPort
	labelingService *labeling.LabelingService
}

func (i *Interactor) Execute(r Request) {
	_, err := i.labelingService.Init(r.ImageId, r.Collection, r.Label)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
		default:
		}
		return
	}
}

func NewInteractor(repo Repo, output OutputPort, labelingService *labeling.LabelingService) *Interactor {
	return &Interactor{repo: repo, output: output, labelingService: labelingService}
}
