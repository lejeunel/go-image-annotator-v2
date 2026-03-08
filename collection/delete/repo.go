package delete

import "slices"

type Repo interface {
	Delete(string) error
	Exists(string) (bool, error)
	IsPopulated(string) (bool, error)
}

type FakeRepo struct {
	Collections  []string
	ArePopulated []string
}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Delete(string) error {
	return r.err
}
func (r *FakeErrRepo) Exists(string) (bool, error) {
	return false, r.err
}
func (r *FakeErrRepo) IsPopulated(string) (bool, error) {
	return false, r.err
}

func (r *FakeRepo) Delete(string) error {
	return nil
}

func (r *FakeRepo) Exists(c string) (bool, error) {
	if slices.Contains(r.Collections, c) {
		return true, nil
	}
	return false, nil
}

func (r *FakeRepo) IsPopulated(c string) (bool, error) {
	if slices.Contains(r.ArePopulated, c) {
		return true, nil
	}
	return false, nil
}
