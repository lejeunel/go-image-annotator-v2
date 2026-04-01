package delete

type FakeRepo struct {
	Err         error
	IsUsed_     bool
	IsMissing   bool
	ErrOnDelete bool
	ErrOnIsUsed bool
	ErrOnExists bool
}

func (r *FakeRepo) Delete(string) error {
	if r.Err != nil {
		return r.Err
	}
	return nil
}

func (r *FakeRepo) IsUsed(n string) (*bool, error) {
	res := true
	if r.ErrOnIsUsed {
		res = false
		return &res, r.Err
	}
	if r.IsUsed_ {
		return &res, nil

	}
	res = false
	return &res, nil
}
func (r *FakeRepo) Exists(n string) (bool, error) {
	if r.ErrOnExists {
		return false, r.Err
	}
	if r.IsMissing {
		return false, nil
	}

	return true, nil

}

type FakePresenter struct {
	GotDependencyErr bool
	GotInternalErr   bool
	GotNotFoundErr   bool
	GotSuccess       bool
}

func (p *FakePresenter) ErrDependency(error) {
	p.GotDependencyErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) Success() {
	p.GotSuccess = true
}
