package create

import (
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	v "github.com/lejeunel/go-image-annotator-v2/shared/validation"
)

func TestCreateCollectionWithDuplicateNameShouldFail(t *testing.T) {
	name := "my-collection"
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Names: []string{name}}, &v.FakeNameValidator{}, clockwork.NewFakeClock())
	itr.Execute(Request{Name: name}, p)
	if !p.GotDuplicationErr {
		t.Fatal("expected duplication error, but go none")
	}
	if p.GotSuccess {
		t.Fatal("expected no success")
	}
}

func TestHandleInternalError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&FakeRepo{Err: e.ErrInternal}, &v.FakeNameValidator{}, clockwork.NewFakeClock())
	itr.Execute(Request{}, p)
	if !p.GotInternalErr {
		t.Fatal("expected internal error, but got none")
	}
}

func TestCreateCollectionWithInvalidNameShouldFail(t *testing.T) {
	name := "my-collection%/"
	p := &FakePresenter{}
	validator := &v.FakeNameValidator{Err: e.ErrValidation}
	itr := NewInteractor(&FakeRepo{Names: []string{name}}, validator, clockwork.NewFakeClock())
	itr.Execute(Request{Name: name}, p)
	if !p.GotValidationErr {
		t.Fatal("expected validation error, but go none")
	}
}

func TestCreateCollection(t *testing.T) {
	p := &FakePresenter{}
	repo := &FakeRepo{}
	now := time.Now()
	itr := NewInteractor(repo, &v.FakeNameValidator{}, clockwork.NewFakeClockAt(now))
	name := "a-name"
	desc := "a-description"
	req := Request{Name: name, Description: desc}
	itr.Execute(req, p)
	got := repo.Got
	if got.Name != name || got.Description != desc || got.Id.IsNil() || got.CreatedAt != now {
		t.Fatalf("expected to create collection with name %v, description %v, non-nil id, and created at %v got %v, %v, %v, %v",
			name, desc, now, got.Name, got.Description, got.Id, got.CreatedAt)

	}
}
