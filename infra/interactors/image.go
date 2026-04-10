package interactors

import (
	has "github.com/lejeunel/go-image-annotator-v2/application/hasher"
	rea "github.com/lejeunel/go-image-annotator-v2/application/image-reader"
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
			repos.ArtefactRepo, has.NewSha256Hasher(), rea.ImageMIMETypeDetector{}),
		Read:                *read.NewInteractor(*repos.ImageStore),
		List:                *list.NewInteractor(repos.ImageRepo, repos.ImageStore),
		AllowedImageFormats: allowedImageFormats,
		DefaultPageSize:     10,
	}
}
