package read

import e "github.com/lejeunel/go-image-annotator-v2/errors"

type ReadCollectionRepo interface {
	Find(string) (*Collection, error)
}

type FakeReadCollectionRepo struct {
	Collection Collection
}

func (r *FakeReadCollectionRepo) Find(name string) (*Collection, error) {

	if name == r.Collection.Name {
		return &r.Collection, nil
	}
	return nil, e.ErrNotFound

}

type FakeInternalErrReadCollectionRepo struct{}

func (r *FakeInternalErrReadCollectionRepo) Find(name string) (*Collection, error) {
	return nil, e.ErrInternal

}
