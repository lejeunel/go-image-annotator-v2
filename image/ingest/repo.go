package ingest

import "slices"

type Repo interface {
	CollectionExists(string) (bool, error)
	LabelExists(string) (bool, error)
}

type FakeRepo struct {
	Collections []string
	Labels      []string
}

func (r *FakeRepo) CollectionExists(name string) (bool, error) {
	if slices.Contains(r.Collections, name) {
		return true, nil
	}
	return false, nil
}

func (r *FakeRepo) LabelExists(name string) (bool, error) {
	if slices.Contains(r.Labels, name) {
		return true, nil
	}
	return false, nil
}

type FakeCollectionExistsErrRepo struct {
	err error
}

func (r *FakeCollectionExistsErrRepo) CollectionExists(string) (bool, error) {
	return false, r.err
}

func (r *FakeCollectionExistsErrRepo) LabelExists(string) (bool, error) {
	return true, nil
}

type FakeLabelExistsErrRepo struct {
	err error
}

func (r *FakeLabelExistsErrRepo) CollectionExists(string) (bool, error) {
	return true, nil
}

func (r *FakeLabelExistsErrRepo) LabelExists(string) (bool, error) {
	return false, r.err
}
