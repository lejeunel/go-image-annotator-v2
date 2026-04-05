package create

import (
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
	"slices"
)

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

type FakeRepo struct {
	Err   error
	Names []string
	Got   lbl.Label
}

func (r *FakeRepo) Create(l lbl.Label) error {
	if r.Err != nil {
		return r.Err
	}
	r.Got = l
	return nil
}
func (r *FakeRepo) Exists(name string) (bool, error) {
	if slices.Contains(r.Names, name) {
		return true, nil
	}
	return false, nil
}
