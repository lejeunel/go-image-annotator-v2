package create

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
)

type CreateRepo interface {
	Create(clc.Collection) error
	Exists(string) (bool, error)
}
