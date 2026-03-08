package errors

import (
	"errors"
)

var ErrInternal = errors.New("internal error")
var ErrDuplicate = errors.New("duplicate resource error")
var ErrNotFound = errors.New("resource not found error")
var ErrDependency = errors.New("dependency error")
var ErrValidation = errors.New("validation error")
