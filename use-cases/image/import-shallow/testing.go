package import_shallow

import (
	clc "github.com/lejeunel/go-image-annotator-v2/domain/collection"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
)

type FakePresenter struct {
	GotNotFoundErr   bool
	GotInternalErr   bool
	GotDependencyErr bool
	GotSuccess       bool
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}
func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrDependency(error) {
	p.GotDependencyErr = true
}

func (p *FakePresenter) Success(Response) {
	p.GotSuccess = true
}

type FakeRepo struct {
	ErrOnImageExists             bool
	ErrOnFindCollection          bool
	ErrOnImageExistsInCollection bool
	ErrOnImport                  bool
	ImageMissing                 bool
	ImageAlreadyInCollection     bool

	ImportedImageId          im.ImageId
	ImportedIntoCollectionId clc.CollectionId

	Err                   error
	DestinationCollection clc.Collection
}

func (r *FakeRepo) ImageExists(imageId im.ImageId) (bool, error) {
	if r.ErrOnImageExists {
		return false, r.Err
	}
	if r.ImageMissing {
		return false, nil
	}
	return true, nil
}

func (r *FakeRepo) FindCollection(name string) (*clc.Collection, error) {
	if r.ErrOnFindCollection {
		return nil, r.Err
	}
	return &r.DestinationCollection, nil
}

func (r *FakeRepo) ImageExistsInCollection(imageId im.ImageId, collectionId clc.CollectionId) (bool, error) {
	if r.ErrOnImageExistsInCollection {
		return false, r.Err
	}
	if r.ImageAlreadyInCollection {
		return true, nil
	}
	return false, nil
}

func (r *FakeRepo) ImportImage(imageId im.ImageId, collectionId clc.CollectionId) error {
	if r.ErrOnImport {
		return r.Err
	}
	r.ImportedImageId = imageId
	r.ImportedIntoCollectionId = collectionId
	return nil
}
