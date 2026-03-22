package labeling

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
)

type FakeLabelingService struct {
	Err error
}

func (s *FakeLabelingService) Init(imageId im.ImageId, collectionName string, labelName string) (*LabelingCtx, error) {
	if s.Err != nil {
		return nil, s.Err
	}
	return &LabelingCtx{ImageId: im.NewImageID(), CollectionId: clc.NewCollectionID(), LabelId: lbl.NewLabelID()}, nil
}
