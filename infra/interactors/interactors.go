package interactors

import (
	"github.com/lejeunel/go-image-annotator-v2/infra"
	u "github.com/lejeunel/go-image-annotator-v2/use-cases"
)

func NewSQLiteInteractors(repos *infra.SQLiteInfra, allowedImageFormats []string) *u.Interactors {

	return &u.Interactors{
		Label:      NewSQLiteLabelInteractors(repos.LabelRepo),
		Collection: NewSQLiteCollectionInteractors(repos.CollectionRepo),
		Image:      NewSQLiteImageInteractors(repos, allowedImageFormats),
	}
}
