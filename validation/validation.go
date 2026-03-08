package validation

import (
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"regexp"
)

type Validator interface {
	Validate(string) error
}

type FakeInvalidNameValidator struct{}

func (v *FakeInvalidNameValidator) Validate(name string) error {
	return e.ErrValidation
}

type FakeValidNameValidator struct{}

func (v *FakeValidNameValidator) Validate(name string) error {
	return nil
}

type NameValidator struct {
}

var validName = regexp.MustCompile(`^[a-z0-9-]+$`)

func (v *NameValidator) Validate(name string) error {
	if !validName.MatchString(name) {
		return e.ErrValidation
	}
	return nil

}

func NewNameValidator() *NameValidator {
	return &NameValidator{}
}
