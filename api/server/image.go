package server

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	presenter "github.com/lejeunel/go-image-annotator-v2/adapters/json/image"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	far "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	has "github.com/lejeunel/go-image-annotator-v2/application/hasher"
	ide "github.com/lejeunel/go-image-annotator-v2/application/image-decoder"
	image_store "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	image "github.com/lejeunel/go-image-annotator-v2/entities/image"
	anr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/annotation"
	clr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	imr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/image"
	lbr "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/ingest"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/read-meta"
)

type ImageServer struct {
	Ingest   ingest.Interactor
	ReadMeta read_meta.Interactor
	List     list.Interactor
}

func NewHTTPImageServer(db *sqlx.DB, baseDir string, allowedImageFormats []string) *ImageServer {
	imRepo := imr.NewSQLiteImageRepo(db)
	clRepo := clr.NewSQLiteCollectionRepo(db)
	lbRepo := lbr.NewSQLiteLabelRepo(db)
	anRepo := anr.NewSQLiteAnnotationRepo(db)
	artRepo := far.NewFileArtefactRepo(baseDir)
	imStore := image_store.NewImageStore(imRepo, clRepo, anRepo, artRepo)
	return &ImageServer{
		Ingest: *ingest.NewInteractor(imRepo, clRepo, lbRepo, anRepo,
			artRepo, has.NewSha256Hasher(), ide.NewBase64ImageDecoder(allowedImageFormats)),
		ReadMeta: *read_meta.NewInteractor(imStore),
		List:     *list.NewInteractor(imRepo, imStore),
	}
}

func (s *Server) IngestImage(w http.ResponseWriter, r *http.Request) {
	body, ok := json.DecodeJSONOrFail[models.NewImage](w, r)
	if !ok {
		return
	}

	req := ingest.Request{Collection: body.Collection, Data: body.Data}
	if body.Labels != nil {
		req.Labels = *body.Labels
	}

	if body.BoundingBoxes != nil {
		for _, bbox := range *body.BoundingBoxes {
			req.BoundingBoxes = append(req.BoundingBoxes,
				ingest.BoundingBoxRequest{Xc: bbox.Xc, Yc: bbox.Yc,
					Width: bbox.Width, Height: bbox.Height})
		}
	}
	s.Image.Ingest.Execute(req,
		&presenter.Ingest{Writer: w})

}

func (s *Server) ReadImage(w http.ResponseWriter, r *http.Request, collectionName, imageId string) {
	id, err := image.NewImageIdFromString(imageId)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("parsing UUID from string: %w", err).Error())
	}
	s.Image.ReadMeta.Execute(read_meta.Request{ImageId: id, Collection: collectionName},
		&presenter.ReadMeta{Writer: w})
}

func (s *Server) ListImages(w http.ResponseWriter, r *http.Request, params ListImagesParams) {
	req := list.Request{Page: 1, PageSize: s.Collection.DefaultPageSize, CollectionName: params.Collection}
	if p := params.Page; p != nil {
		req.Page = *p
	}
	if p := params.PageSize; p != nil {
		req.PageSize = *p
	}
	s.Image.List.Execute(req, &presenter.List{Writer: w})
}
