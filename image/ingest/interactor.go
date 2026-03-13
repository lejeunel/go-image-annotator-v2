package ingest

import (
	"fmt"
	"io"

	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

var errCtx = "ingesting image"

type Interactor struct {
	output       OutputPort
	repo         Repo
	artefactRepo ArtefactRepo
}

func NewInteractor(repo Repo, artefactRepo ArtefactRepo, output OutputPort) *Interactor {
	return &Interactor{output: output, repo: repo, artefactRepo: artefactRepo}
}

func (i *Interactor) Execute(r Request) {
	if !i.collectionExists(r.Collection) {
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

	if err := i.artefactRepo.Store(data); err != nil {
		i.output.ErrInternal(fmt.Errorf("%v: storing image data in artefact repository: %w", errCtx, e.ErrInternal))
		return
	}

	i.output.Success(Response{Collection: r.Collection})

}

func (i *Interactor) collectionExists(name string) bool {
	collectionExists, err := i.repo.CollectionExists(name)
	baseErrMsg := fmt.Sprintf("%v: checking whether collection %v exists", errCtx, name)
	if err != nil {
		i.output.ErrInternal(fmt.Errorf("%v: %w", baseErrMsg, e.ErrInternal))
		return false
	}
	if !collectionExists {
		i.output.ErrCollectionNotFound(fmt.Errorf("%v: %w", baseErrMsg, e.ErrNotFound))
		return false
	}
	return true

}

func (i *Interactor) labelsExist(labels []string) bool {
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
