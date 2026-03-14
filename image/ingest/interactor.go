package ingest

import (
	"errors"
	"fmt"
	"io"

	e "github.com/lejeunel/go-image-annotator-v2/errors"

	a "github.com/lejeunel/go-image-annotator-v2/domain/artefact"
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

var errCtx = "ingesting image"

type Interactor struct {
	output       OutputPort
	repo         Repo
	artefactRepo a.ArtefactRepo
}

func NewInteractor(repo Repo, artefactRepo a.ArtefactRepo, output OutputPort) *Interactor {
	return &Interactor{output: output, repo: repo, artefactRepo: artefactRepo}
}

func (i *Interactor) Execute(r Request) {
	collection := i.findCollectionByName(r.Collection)
	if collection == nil {
		return
	}

	if !i.labelsExist(r.Labels) {
		return
	}
	data, err := io.ReadAll(r.Reader)
	if err != nil {
		i.output.ErrInvalidImageData(fmt.Errorf("%v: reading image data: %w", errCtx, e.ErrValidation))
		return
	}

	artefactID := a.NewArtefactID()
	if err := i.artefactRepo.Store(artefactID, data); err != nil {
		i.output.ErrInternal(fmt.Errorf("%v: storing image data in artefact repository: %w", errCtx, e.ErrInternal))
		return
	}

	imageID := im.NewImageID()
	if err := i.repo.IngestImage(imageID, collection.ID, artefactID); err != nil {
		i.output.Success(Response{Collection: r.Collection})
	}
	i.output.ErrInternal(fmt.Errorf("%v: ingesting image meta-data: %w", errCtx, e.ErrInternal))

}

func (i *Interactor) findCollectionByName(name string) *clc.Collection {
	collection, err := i.repo.FindCollectionByName(name)
	baseErrMsg := fmt.Sprintf("%v: finding collection with name %v", errCtx, name)
	switch {
	case errors.Is(err, e.ErrNotFound):
		i.output.ErrCollectionNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
		return nil
	case errors.Is(err, e.ErrInternal):
		i.output.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return nil
	}
	return collection

}

func (i *Interactor) labelsExist(labels []string) bool {
	if len(labels) == 0 {
		return true
	}

	for _, label := range labels {
		baseErrMsg := fmt.Sprintf("%v: checking whether label %v exist", errCtx, label)
		labelExists, err := i.repo.LabelExists(label)
		if err != nil {
			i.output.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
			return false
		}
		if !labelExists {
			i.output.ErrLabelNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
			return false
		}
	}
	return true
}
