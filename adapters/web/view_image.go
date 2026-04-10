package web

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"text/template"

	"embed"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	n "github.com/lejeunel/go-image-annotator-v2/shared/navigation"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/read"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed templates/*
var templatesFiles embed.FS

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
		html.NewPageBuilder().SetError(err.Error()).Render(p.Writer)
		return

	}
	b64Image := base64.StdEncoding.EncodeToString(bytes)
	table := html.SpecTable{}
	table.Rows = append(table.Rows, html.SpecTableRow{Name: "id", Value: image.Id.String()})
	table.Rows = append(table.Rows, html.SpecTableRow{Name: "collection", Value: image.Collection.Name})
	imNode := Img(ID("image"), Src(fmt.Sprintf("data:%v;base64,%s", image.MIMEType, b64Image)))
	tableNode := table.Render()

	b := html.NewTitledPageBuilder("Image")
	b.AddScripts(html.AnnotoriousScripts()...)
	s, err := p.makeAnnotationScript(image)
	if err != nil {
		html.NewPageBuilder().SetError(err.Error()).Render(p.Writer)
		return
	}
	b.AddScripts(Script(Raw(s)))
	b.SetContent(Table(Tr(Td(imNode), Td(Class("align-top pl-2"), tableNode))))
	b.Render(p.Writer)
}

type AnnotatorState struct {
	ImageId          string
	Collection       string
	Annotations      string
	EnableAnnotation bool
	OriginType       string
	OriginId         string
	Ordering         string
	Descending       bool
}

func (p ViewImagePresenter) makeAnnotationScript(image *im.Image) (string, error) {
	tAnnot, err := template.New("annotator").ParseFS(templatesFiles, "templates/annotator.js")
	if err != nil {
		return "", err
	}
	buf := bytes.NewBufferString("")
	data := AnnotatorState{ImageId: image.Id.String(),
		Collection: image.Collection.Name, Annotations: "[]",
		EnableAnnotation: true}

	err = tAnnot.ExecuteTemplate(buf, "annotator", data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *Server) ViewImage(w http.ResponseWriter, r *http.Request) {

	req, err := ParseImageURL(r.URL)
	if err != nil {
		html.NewPageBuilder().SetError(err.Error()).Render(w)
		return
	}
	s.Image.Read.Execute(*req, NewViewImagePresenter(w))

}

func NewViewImagePresenter(w http.ResponseWriter) ViewImagePresenter {
	baseURL, _ := url.Parse("/collections")
	return ViewImagePresenter{
		ListRenderer: NewListRenderer("Collections", *baseURL,
			n.CollectionsPageActive, w),
	}
}
