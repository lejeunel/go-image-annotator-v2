package update

import (
	"slices"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type FakeRepo struct {
	Names []string
	Got   Model
}

func (r *FakeRepo) Update(m Model) error {
	r.Got = m
	return nil
}
func (r *FakeRepo) Exists(n string) (bool, error) {
	if slices.Contains(r.Names, n) {
		return true, nil
	}
	return false, nil
}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Update(m Model) error {
	return e.ErrInternal
}

func (r *FakeErrRepo) Exists(n string) (bool, error) {
	return false, r.err
}

type FakePresenter struct {
	Got               Response
	GotDuplicationErr bool
	GotNotFoundErr    bool
	GotInternalErr    bool
	GotSuccess        bool
}

func (p *FakePresenter) ErrDuplication(error) {
	p.GotDuplicationErr = true
}

func (p *FakePresenter) ErrNotFound(error) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(error) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
