package interactors

import (
	i "github.com/lejeunel/go-image-annotator-v2/application/interactors"
	"github.com/lejeunel/go-image-annotator-v2/infra"
)

func NewSQLiteInteractors(repos *infra.SQLiteInfra, allowedImageFormats []string) *i.Interactors {

	return &i.Interactors{
		Label:      NewSQLiteLabelInteractors(repos.LabelRepo),
		Collection: NewSQLiteCollectionInteractors(repos.CollectionRepo),
		Image:      NewSQLiteImageInteractors(repos, allowedImageFormats),
	}
}
