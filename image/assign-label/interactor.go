package assign_label

import (
	"errors"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	labeling "github.com/lejeunel/go-image-annotator-v2/domain/labeling"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo            Repo
	output          OutputPort
	labelingService *labeling.LabelingService
}

func (i *Interactor) Execute(r Request) {
	ctx, err := i.labelingService.Init(r.ImageId, r.Collection, r.Label)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
		case errors.Is(err, e.ErrImageNotInCollection):
			i.output.ErrImageNotInCollection(err)
		default:
			i.output.ErrInternal(err)
		}
		return
	}

	if ok := i.addLabel(r.ImageId, ctx.Collection.Id, ctx.Label.Id); !ok {
		return
	}

	i.output.Success(Response{ImageId: r.ImageId, Collection: ctx.Collection.Name, Label: ctx.Label.Name})
}

func (i *Interactor) addLabel(imageId im.ImageId, collectionId clc.CollectionId, labelId lbl.LabelId) bool {
	err := i.repo.AddLabel(imageId, collectionId, labelId)
	if err != nil {
		i.output.ErrInternal(err)
		return false
	}
	return true

}

func NewInteractor(repo Repo, output OutputPort, labelingService *labeling.LabelingService) *Interactor {
	return &Interactor{repo: repo, output: output, labelingService: labelingService}
}
