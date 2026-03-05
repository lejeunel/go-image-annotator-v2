package collection

import (
	"errors"
)

var ErrDuplicate = errors.New("duplicate collection error")
var ErrNotFound = errors.New("collection not found error")
