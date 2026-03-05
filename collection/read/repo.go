package read

import c "github.com/lejeunel/go-image-annotator-v2/collection"

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
	return nil, c.ErrNotFound

}
