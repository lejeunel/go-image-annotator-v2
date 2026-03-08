package delete

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"slices"
)

type FakeRepo struct {
	ArePopulated []string
}

type FakeErrRepo struct {
	err error
}

func (r *FakeErrRepo) Delete(m Model) error {
	return e.ErrInternal
}

func (r *FakeRepo) Delete(m Model) error {
	if slices.Contains(r.ArePopulated, m.Name) {
		return e.ErrDependency
	}

	return nil
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
