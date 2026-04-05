package read

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
)

type Repo interface {
	FindCollectionByName(string) (*clc.Collection, error)
}
