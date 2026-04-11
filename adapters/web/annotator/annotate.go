package annotator

import (
	"io"

	"embed"

	scr "github.com/lejeunel/go-image-annotator-v2/application/scroller"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	. "maragu.dev/gomponents/html"
)

//go:embed templates/*
var templatesFiles embed.FS

type AnnotationView struct {
	ImageView      ImageView
	ImageInfosView ImageInfosView
	ScrollerView   ScrollerView
	image          im.Image
	scroller       scr.ScrollerState
	err            error
}

func (p *AnnotationView) RenderError(err error, w io.Writer) {
	b := html.NewTitledPageBuilder("Image")
	b.SetError(err).Render(w)
}
func (v *AnnotationView) DrawScroller(scroller scr.ScrollerState) {
	v.scroller = scroller
}

func (v *AnnotationView) DrawImage(image im.Image) {
	v.image = image
}
func (v *AnnotationView) Error(err error) {
	v.err = err
}

func (v *AnnotationView) Render(w io.Writer) {

	if v.err != nil {
		html.NewPageBuilder().SetError(v.err).Render(w)
	}

	b := html.NewTitledPageBuilder("Image")
	script, err := MakeAnnotoriousScript(v.image.Id, v.image.Collection.Name)
	if err != nil {
		b.SetError(err).Render(w)
		return
	}
	b.AddScripts(html.AnnotoriousLib()...)
	b.AddScripts(*script)
	b.SetContent(
		Table(
			Tr(Td(v.ScrollerView.Render(v.scroller))),
			Tr(Td(Table(
				Tr(Td(v.ImageView.Render(&v.image)),
					Td(Class("align-top pl-2"), v.ImageInfosView.Render(&v.image)))),
			))))
	b.Render(w)

}

func NewAnnotationView() *AnnotationView {
	return &AnnotationView{
		ImageView:      ImageView{},
		ImageInfosView: ImageInfosView{},
		ScrollerView:   ScrollerView{},
	}
}
