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
	Id         im.ImageId
	Collection string
}

func ParseImageURL(u *url.URL) (*Request, error) {
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
	req.Id = imageId

	collection := u.Query().Get("collection")
	if collection == "" {
		return nil, fmt.Errorf("%v: collection (%v): %w", baseErr, collection, e.ErrURLParsing)
	}
	req.Collection = collection
	return &req, nil
}

func (s *Server) AnnotateImage(w http.ResponseWriter, r *http.Request) {

	p := aw.NewAnnotationView()
	req, err := ParseImageURL(r.URL)
	if err != nil {
		p.RenderError(err, w)
		return
	}

	annotator, err := s.annotatorBuilder.Build(req.Id, req.Collection)
	if err != nil {
		p.RenderError(err, w)
		return
	}
	state, err := annotator.State()
	if err != nil {
		p.RenderError(err, w)
		return
	}
	p.Render(*state, w)

}
