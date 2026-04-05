package list

import (
	l "github.com/lejeunel/go-image-annotator-v2/entities/label"
)

type Repo interface {
	List(Request) ([]*l.Label, error)
	Count() (int64, error)
}
