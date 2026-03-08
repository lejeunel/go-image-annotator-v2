package delete

import "slices"

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Delete(string) error {
	return r.err
}
func (r *FakeErrRepo) IsUsed(string) (bool, error) {
	return false, r.err
}

type FakeRepo struct {
	Used []string
}

func (r *FakeRepo) Delete(string) error {
	return nil
}

func (r *FakeRepo) IsUsed(n string) (bool, error) {
	if slices.Contains(r.Used, n) {
		return true, nil
	}
	return false, nil
}

type FakePresenter struct {
	GotDependencyErr bool
	GotInternalErr   bool
	GotSuccess       bool
}

func (p *FakePresenter) ErrDependency(error) {
	p.GotDependencyErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success() {
	p.GotSuccess = true
}
