package interactors

import (
	has "github.com/lejeunel/go-image-annotator-v2/application/hasher"
	i "github.com/lejeunel/go-image-annotator-v2/application/interactors"
	"github.com/lejeunel/go-image-annotator-v2/infra"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/ingest"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/read"
)

func NewSQLiteImageInteractors(repos *infra.SQLiteInfra, allowedImageFormats []string) *i.ImageInteractors {
	return &i.ImageInteractors{
		Ingest: *ingest.NewInteractor(repos.ImageRepo, repos.CollectionRepo,
			repos.LabelRepo, repos.AnnotationRepo,
			repos.ArtefactRepo, has.NewSha256Hasher()),
		Read:                *read.NewInteractor(*repos.ImageStore),
		List:                *list.NewInteractor(repos.ImageRepo, repos.ImageStore),
		AllowedImageFormats: allowedImageFormats,
		DefaultPageSize:     10,
	}
}
