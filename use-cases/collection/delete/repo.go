package delete

type Repo interface {
	Delete(string) error
	Exists(string) (bool, error)
	IsPopulated(string) (*bool, error)
}

type FakeRepo struct {
	Err          error
	ErrOnDelete  bool
	Missing      bool
	IsPopulated_ bool
}

func (r *FakeRepo) Delete(string) error {

	if r.ErrOnDelete {
		return r.Err
	}
	return nil
}

func (r *FakeRepo) Exists(c string) (bool, error) {
	if r.Missing {
		return false, nil
	}
	return true, nil
}

func (r *FakeRepo) IsPopulated(c string) (*bool, error) {
	res := true
	if r.IsPopulated_ {
		return &res, nil
	}
	res = false
	return &res, nil
}
