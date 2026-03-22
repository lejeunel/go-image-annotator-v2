package unassign_label

import (
	"errors"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	labeling "github.com/lejeunel/go-image-annotator-v2/domain/labeling"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo             Repo
	output           OutputPort
	labelCtxProvider labeling.LabelContextProvider
}

func (i *Interactor) Execute(r Request) {
	ctx, err := i.labelCtxProvider.Init(r.ImageId, r.Collection, r.Label)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
		case errors.Is(err, e.ErrDependency):
			i.output.ErrDependency(err)
		default:
			i.output.ErrInternal(err)
		}
		return
	}
	if ok := i.removeLabel(r.ImageId, ctx.CollectionId, ctx.LabelId); !ok {
		return
	}
	i.output.Success(Response{})
}

func (i *Interactor) removeLabel(imageId im.ImageId, collectionId clc.CollectionId, labelId lbl.LabelId) bool {

	err := i.repo.RemoveLabel(imageId, collectionId, labelId)
	if err != nil {
		i.output.ErrInternal(err)
		return false
	}
	return true
}

func NewInteractor(repo Repo, output OutputPort, labelingService labeling.LabelContextProvider) *Interactor {
	return &Interactor{repo: repo, output: output, labelCtxProvider: labelingService}
}
