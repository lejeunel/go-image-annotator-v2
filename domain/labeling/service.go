package labeling

import (
	"fmt"

	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type LabelingService struct {
	repo Repo
}

type LabelingCtx struct {
	ImageId    im.ImageId
	Collection *clc.Collection
	Label      *lbl.Label
}

func (s *LabelingService) Init(imageId im.ImageId, collectionName string, labelName string) (*LabelingCtx, error) {
	errCtx := "initializing labeling context"
	label, err := s.repo.FindLabelByName(labelName)
	if err != nil {
		return nil, fmt.Errorf("%v: finding label %v: %w", errCtx, labelName, err)
	}
	collection, err := s.repo.FindCollectionByName(collectionName)
	if err != nil {
		return nil, fmt.Errorf("%v: finding collection %v: %w", errCtx, collection, err)
	}
	imageInCollection, err := s.repo.ImageIsInCollection(imageId, collection.Id)
	if err != nil {
		return nil, fmt.Errorf("%v: checking whether image %v is member of collection %v: %w", errCtx, imageId, collectionName, e.ErrInternal)
	}
	if !imageInCollection {
		return nil, fmt.Errorf("%v: checking whether image %v is member of collection %v: %w", errCtx, imageId, collectionName, e.ErrImageNotInCollection)
	}
	return &LabelingCtx{ImageId: imageId, Collection: collection, Label: label}, nil

}

func NewLabelingService(repo Repo) *LabelingService {
	return &LabelingService{repo: repo}
}
