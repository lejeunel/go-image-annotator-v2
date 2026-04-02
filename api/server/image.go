package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	presenter "github.com/lejeunel/go-image-annotator-v2/adapters/json/image"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	far "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	has "github.com/lejeunel/go-image-annotator-v2/application/hasher"
	ide "github.com/lejeunel/go-image-annotator-v2/application/image-decoder"
	anr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/annotation"
	clr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	imr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/image"
	lbr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/ingest"
	"net/http"
)

type ImageServer struct {
	Ingest ingest.Interactor
}

func NewHTTPImageServer(db *sqlx.DB, baseDir string, allowedImageFormats []string) *ImageServer {
	imRepo := imr.NewSQLiteImageRepo(db)
	clRepo := clr.NewSQLiteCollectionRepo(db)
	lbRepo := lbr.NewSQLiteLabelRepo(db)
	anRepo := anr.NewSQLiteAnnotationRepo(db)
	return &ImageServer{
		Ingest: *ingest.NewInteractor(imRepo, clRepo, lbRepo, anRepo,
			far.NewFileArtefactRepo(baseDir), has.NewSha256Hasher(), ide.NewBase64ImageDecoder(allowedImageFormats)),
	}
}

func (s *Server) IngestImage(w http.ResponseWriter, r *http.Request) {
	body, ok := json.DecodeJSONOrFail[models.NewImage](w, r)
	if !ok {
		return
	}

	s.Image.Ingest.Execute(ingest.Request{Collection: body.Collection, Data: body.Data},
		&presenter.Ingest{Writer: w})

}
