package annotator

import (
	"io"

	"embed"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	. "maragu.dev/gomponents/html"
)

//go:embed templates/*
var templatesFiles embed.FS

type AnnotationView struct {
	ImageView      ImageView
	ImageInfosView ImageInfosView
}

func (p *AnnotationView) Render(imageId im.ImageId, collection string, w io.Writer) {

	b := html.NewTitledPageBuilder("Image")
	b.AddScripts(html.AnnotoriousLib()...)
	script, err := MakeAnnotoriousScript(imageId, collection)
	if err != nil {
		b.SetError(err.Error()).Render(w)
		return
	}
	b.AddScripts(*script)
	b.SetContent(Table(Tr(Td(p.ImageView.Build()),
		Td(Class("align-top pl-2"), p.ImageInfosView.Build()))))
	b.Render(w)

}

func NewAnnotationView() *AnnotationView {
	return &AnnotationView{
		ImageView:      ImageView{},
		ImageInfosView: ImageInfosView{},
	}
}
