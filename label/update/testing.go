package update

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"slices"
)

type FakeRepo struct {
	Names []string
	Got   Model
}

func (r *FakeRepo) Update(m Model) error {
	if !slices.Contains(r.Names, m.Name) {
		return e.ErrNotFound
	}
	if slices.Contains(r.Names, m.NewName) {
		return e.ErrDuplicate
	}
	r.Got = m
	return nil
}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Update(m Model) error {
	return r.err
}

type FakePresenter struct {
	Got               Response
	GotDuplicationErr bool
	GotNotFoundErr    bool
	GotInternalErr    bool
	GotSuccess        bool
}

func (p *FakePresenter) ErrDuplication(m string) {
	p.GotDuplicationErr = true
}

func (p *FakePresenter) ErrNotFound(m string) {
	p.GotNotFoundErr = true
}

func (p *FakePresenter) ErrInternal(m string) {
	p.GotInternalErr = true
}

func (p *FakePresenter) Success(r Response) {
	p.GotSuccess = true
	p.Got = r
}
