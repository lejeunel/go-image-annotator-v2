package create

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	"slices"
)

type FakeRepo struct {
	Err   error
	Names []string
	Got   clc.Collection
}

func (r *FakeRepo) Create(c clc.Collection) error {
	if r.Err != nil {
		return r.Err
	}

	r.Got = c
	return nil
}

func (r *FakeRepo) Exists(name string) (bool, error) {
	if slices.Contains(r.Names, name) {
		return true, nil
	}
	return false, nil
}

type FakePresenter struct {
	Got               Response
	GotSuccess        bool
	GotDuplicationErr bool
	GotInternalErr    bool
	GotValidationErr  bool
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
func (p *FakePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}
func (p *FakePresenter) ErrValidation(error) {
	p.GotValidationErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}
