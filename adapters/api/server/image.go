package server

import (
	"fmt"
	"net/http"

	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	presenter "github.com/lejeunel/go-image-annotator-v2/adapters/api/json/image"
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/models"
	app "github.com/lejeunel/go-image-annotator-v2/application"
	has "github.com/lejeunel/go-image-annotator-v2/application/hasher"
	dec "github.com/lejeunel/go-image-annotator-v2/application/image-decoder"
	image "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/ingest"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/read-meta"
)

type ImageServer struct {
	Ingest              ingest.Interactor
	ReadMeta            read_meta.Interactor
	List                list.Interactor
	AllowedImageFormats []string
}

func NewHTTPImageServer(app *app.SQLiteApp, allowedImageFormats []string) *ImageServer {
	return &ImageServer{
		Ingest: *ingest.NewInteractor(app.ImageRepo, app.CollectionRepo,
			app.LabelRepo, app.AnnotationRepo,
			app.ArtefactRepo, has.NewSha256Hasher()),
		ReadMeta:            *read_meta.NewInteractor(*app.ImageStore),
		List:                *list.NewInteractor(app.ImageRepo, app.ImageStore),
		AllowedImageFormats: allowedImageFormats,
	}
}

func (s *Server) IngestImage(w http.ResponseWriter, r *http.Request) {
	body, ok := json.MustDecodeJSON[models.NewImage](w, r)
	if !ok {
		return
	}

	s.Image.Ingest.Execute(NewIngestImageRequest(*body, s.Image.AllowedImageFormats),
		presenter.NewIngestPresenter(w))
}

func (s *Server) ReadImage(w http.ResponseWriter, r *http.Request, collectionName, imageId string) {
	id, err := image.NewImageIdFromString(imageId)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("parsing UUID from string: %w", err).Error())
	}
	s.Image.ReadMeta.Execute(read_meta.Request{ImageId: id, Collection: collectionName},
		presenter.NewReadMetaPresenter(w))
}

func (s *Server) ListImages(w http.ResponseWriter, r *http.Request, params ListImagesParams) {
	req := list.Request{Page: 1, PageSize: s.Collection.DefaultPageSize, CollectionName: params.Collection}
	if p := params.Page; p != nil {
		req.Page = *p
	}
	if p := params.PageSize; p != nil {
		req.PageSize = *p
	}
	s.Image.List.Execute(req, presenter.NewListPresenter(w))
}

func NewIngestImageRequest(req models.NewImage, allowedImageFormats []string) ingest.Request {

	ingestReq := ingest.Request{Collection: req.Collection, Reader: dec.NewBase64ImageDecoder(allowedImageFormats, req.Data)}
	appendLabelsToIngestImageRequest(&ingestReq, req.Labels)
	appendBoundingBoxesToIngestImageRequest(&ingestReq, req.BoundingBoxes)
	return ingestReq
}

func appendBoundingBoxesToIngestImageRequest(req *ingest.Request, boxes *[]models.NewBoundingBox) {
	if boxes != nil {
		for _, box := range *boxes {
			req.BoundingBoxes = append(req.BoundingBoxes,
				ingest.BoundingBoxRequest{Xc: box.Xc, Yc: box.Yc,
					Width: box.Width, Height: box.Height})
		}
	}
}

func appendLabelsToIngestImageRequest(req *ingest.Request, labels *[]string) {
	if labels != nil {
		req.Labels = *labels
	}
}
