package list

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
)

type Repo interface {
	List(Request) ([]*clc.Collection, error)
	Count() (*int64, error)
}
