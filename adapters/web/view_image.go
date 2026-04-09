package web

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/read"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
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

type ViewImagePresenter struct {
	ListRenderer
}

func (p ViewImagePresenter) Success(image *im.Image) {
	bytes, err := io.ReadAll(image.Reader)

	if err != nil {
		html.MakeErrorPage(err.Error())
		return

	}
	b64Image := base64.StdEncoding.EncodeToString(bytes)

	content := Div(Text(image.Id.String()), Text(image.Collection.Name),
		Img(Src(fmt.Sprintf("data:image/jpg;base64,%s", b64Image))))
	html.MakeBasePage("Image", content,
		html.Scripts(html.ScriptIncludes{}),
		html.NavBarActivatedItems{}).Render(p.Writer)
}

func (s *Server) ViewImage(w http.ResponseWriter, r *http.Request) {

	req, err := ParseImageURL(r.URL)
	if err != nil {
		html.MakeErrorPage(err.Error())
	}
	s.Image.Read.Execute(*req, NewViewImagePresenter(w))

}

func NewViewImagePresenter(w http.ResponseWriter) ViewImagePresenter {
	baseURL, _ := url.Parse("/collections")
	return ViewImagePresenter{
		ListRenderer: NewListRenderer("Collections", *baseURL,
			html.NavBarActivatedItems{Collections: true}, w),
	}
}
