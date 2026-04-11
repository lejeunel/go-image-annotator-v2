package annotator

import (
	"io"

	"embed"

	a "github.com/lejeunel/go-image-annotator-v2/application/annotator"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	. "maragu.dev/gomponents/html"
)

//go:embed templates/*
var templatesFiles embed.FS

type AnnotationView struct {
	ImageView      ImageView
	ImageInfosView ImageInfosView
	ScrollerView   ScrollerView
}

func (p *AnnotationView) RenderError(err error, w io.Writer) {
	b := html.NewTitledPageBuilder("Image")
	b.SetError(err).Render(w)
}
func (p *AnnotationView) Render(s a.AnnotatorState, w io.Writer) {

	b := html.NewTitledPageBuilder("Image")
	script, err := MakeAnnotoriousScript(s.Image.Id, s.Image.Collection.Name)
	if err != nil {
		b.SetError(err).Render(w)
		return
	}
	b.AddScripts(html.AnnotoriousLib()...)
	b.AddScripts(*script)
	b.SetContent(
		Table(
			Tr(Td(p.ScrollerView.Render(s.Scroller))),
			Tr(Td(Table(
				Tr(Td(p.ImageView.Render(&s.Image)),
					Td(Class("align-top pl-2"), p.ImageInfosView.Render(&s.Image)))),
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
