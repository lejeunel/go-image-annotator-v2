package web

import (
	"fmt"
	a "github.com/lejeunel/go-image-annotator-v2/adapters/web/annotator"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/read"
	"net/http"
	"net/url"
)

func ParseImageURL(u *url.URL) (*read.Request, error) {
	baseErr := "parsing url"
	req := read.Request{}
	imageIdStr := u.Query().Get("id")
	if imageIdStr == "" {
		return nil, fmt.Errorf("%v: extracting id: %w", baseErr, e.ErrURLParsing)
	}
	imageId, err := im.NewImageIdFromString(imageIdStr)
	if err != nil {
		return nil, fmt.Errorf("%v: validating id (%v): %w", baseErr, imageIdStr, e.ErrValidation)

	}
	req.ImageId = imageId

	collection := u.Query().Get("collection")
	if collection == "" {
		return nil, fmt.Errorf("%v: collection (%v): %w", baseErr, collection, e.ErrURLParsing)
	}
	req.Collection = collection
	return &req, nil
}

func (s *Server) AnnotateImage(w http.ResponseWriter, r *http.Request) {

	req, err := ParseImageURL(r.URL)
	if err != nil {
		html.NewPageBuilder().SetError(err.Error()).Render(w)
		return
	}
	p := a.NewAnnotationView()
	s.Image.Read.Execute(*req, &p.ImageView)
	p.Render(req.ImageId, req.Collection, w)

}
