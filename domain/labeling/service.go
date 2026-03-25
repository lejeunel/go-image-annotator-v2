package labeling

import (
	"fmt"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type LabelingCtx struct {
	ImageId      im.ImageId
	CollectionId clc.CollectionId
	LabelId      lbl.LabelId
}

type LabelingCtxInitializer interface {
	Init(imageId im.ImageId, collectionName string, labelName string) (*LabelingCtx, error)
}

type MyLabelingCtxProvider struct {
	repo Repo
}

func (s *MyLabelingCtxProvider) Init(imageId im.ImageId, collectionName string, labelName string) (*LabelingCtx, error) {
	errCtx := "initializing labeling context"
	labelId, err := s.repo.FindLabelIdByName(labelName)
	if err != nil {
		return nil, fmt.Errorf("%v: finding label %v: %w", errCtx, labelName, err)
	}
	collectionId, err := s.repo.FindCollectionIdByName(collectionName)
	if err != nil {
		return nil, fmt.Errorf("%v: finding collection %v: %w", errCtx, collectionName, err)
	}
	if err := s.imageIsInCollection(imageId, *collectionId); err != nil {
		return nil, fmt.Errorf("%v: checking whether image is in collection: %w", errCtx, err)
	}
	return &LabelingCtx{ImageId: imageId, CollectionId: *collectionId, LabelId: *labelId}, nil
}

func (s *MyLabelingCtxProvider) imageIsInCollection(imageId im.ImageId, collectionId clc.CollectionId) error {
	imageInCollection, err := s.repo.ImageIsInCollection(imageId, collectionId)
	if err != nil {
		return e.ErrInternal
	}
	if !imageInCollection {
		return fmt.Errorf("image not member of collection: %w", e.ErrDependency)
	}
	return nil

}

func NewLabelingService(repo Repo) *MyLabelingCtxProvider {
	return &MyLabelingCtxProvider{repo: repo}
}
