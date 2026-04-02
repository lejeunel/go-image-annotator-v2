package ingest

import (
	an "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type FakeHasher struct {
	Hash_ string
}

func (h *FakeHasher) Hash([]byte) string {
	return h.Hash_

}

type FakePresenter struct {
	Got               *Response
	GotSuccess        bool
	GotNotFoundErr    bool
	GotDuplicateImage bool
	GotInternalErr    bool
	GotValidationErr  bool
}

func (p *FakePresenter) Success(r Response) {
	p.Got = &r
	p.GotSuccess = true
}
func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}

func (p *FakePresenter) ErrDuplication(error) {
	p.GotDuplicateImage = true
}

type FakeCollectionRepo struct {
	Err                 error
	ErrOnFindCollection bool
	MissingCollection   bool
}

type FakeLabelRepo struct {
	Err              error
	ErrOnLabelExists bool
	MissingLabel     bool
}

type FakeAnnotationRepo struct {
	Err                   error
	ErrOnAddBoundingBox   bool
	ErrOnAddLabel         bool
	NumLabelsAdded        int
	NumBoundingboxesAdded int
}

type FakeImageRepo struct {
	Err                       error
	GotImage                  bool
	GotHash                   string
	ErrOnAddImageToCollection bool
	ErrOnAddImage             bool
	ErrOnFindHash             bool
	ErrOnDeleteImage          bool
	HashAlreadyExists         bool
	NumDeletedImages          int
}

func (r *FakeCollectionRepo) FindCollectionByName(name string) (*clc.Collection, error) {
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	if r.MissingCollection {
		return nil, e.ErrNotFound
	}
	c := clc.NewCollection(clc.NewCollectionId(), "a-collection")
	return c, nil
}

func (r *FakeLabelRepo) FindLabelByName(name string) (*lbl.Label, error) {
	if r.ErrOnLabelExists {
		return nil, r.Err
	}
	if r.MissingLabel {
		return nil, e.ErrNotFound
	}
	return lbl.NewLabel(lbl.NewLabelId(), name), nil
}

func (r *FakeImageRepo) FindImageIdByHash(hash string) (*im.ImageId, error) {
	if r.ErrOnFindHash {
		return nil, r.Err
	}
	if r.HashAlreadyExists {
		existingId := im.NewImageId()
		return &existingId, nil
	}
	return nil, e.ErrNotFound
}

func (r *FakeAnnotationRepo) AddImageLabel(an.AnnotationId, im.ImageId, clc.CollectionId, lbl.LabelId) error {
	if r.ErrOnAddLabel {
		return r.Err
	}
	r.NumLabelsAdded += 1
	return nil
}

func (r *FakeAnnotationRepo) AddBoundingBox(im.ImageId, clc.CollectionId, an.BoundingBox) error {
	if r.ErrOnAddBoundingBox {
		return r.Err
	}
	r.NumBoundingboxesAdded += 1
	return nil
}

func (r *FakeImageRepo) Delete(im.ImageId) error {
	if r.ErrOnDeleteImage {
		return r.Err
	}
	r.NumDeletedImages += 1
	return nil
}

func (r *FakeImageRepo) AddImageToCollection(im.ImageId, clc.CollectionId) error {
	if r.ErrOnAddImageToCollection {
		return r.Err
	}
	return nil
}

func (r *FakeImageRepo) AddImage(imageId im.ImageId, hash string) error {
	if r.ErrOnAddImage {
		return r.Err
	}
	r.GotHash = hash
	return nil
}

type FakeImageDecoder struct {
	Format_ string
	Err     error
}

func (d *FakeImageDecoder) Decode(data any) ([]byte, *string, error) {
	if d.Err != nil {
		return nil, nil, d.Err
	}
	return nil, &d.Format_, nil

}
