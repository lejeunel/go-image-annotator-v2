package web

import (
	"fmt"
	aw "github.com/lejeunel/go-image-annotator-v2/adapters/web/annotator"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	"net/http"
	"net/url"
)

type Request struct {
	ImageId    im.ImageId
	Collection string
}

func ParseURL(u *url.URL) (*Request, error) {
	baseErr := "parsing url"
	req := Request{}
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

func (s *Server) ViewImage(w http.ResponseWriter, r *http.Request) {

	view := aw.NewAnnotationView()
	req, err := ParseURL(r.URL)
	if err != nil {
		view.RenderError(err, w)
		return
	}

	annotator, err := s.annotatorBuilder.Build(req.ImageId, req.Collection)
	if err != nil {
		view.RenderError(err, w)
		return
	}
	state, err := annotator.State()
	if err != nil {
		view.RenderError(err, w)
		return
	}
	view.Render(*state, w)

}
