package create

import (
	"slices"
)

type FakeRepo struct {
	Names []string
	Got   Model
}

func (r *FakeRepo) Create(m Model) error {
	r.Got = m
	return nil
}
func (r *FakeRepo) Exists(name string) (bool, error) {
	if slices.Contains(r.Names, name) {
		return true, nil
	}
	return false, nil
}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Create(m Model) error {
	return r.err
}
func (r *FakeErrRepo) Exists(name string) (bool, error) {
	return false, r.err
}

type FakePresenter struct {
	Got               Response
	GotDuplicationErr bool
	GotInternalErr    bool
	GotSuccess        bool
	GotValidationErr  bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
func (p *FakePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}
func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}
