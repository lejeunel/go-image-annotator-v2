package annotator

import (
	"io"

	"embed"

	a "github.com/lejeunel/go-image-annotator-v2/application/annotator"
	scr "github.com/lejeunel/go-image-annotator-v2/application/scroller"
	an "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	addbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/add-bbox"
	updbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/modify-bbox"
	del "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/remove"
	. "maragu.dev/gomponents/html"
)

//go:embed templates/*
var templatesFiles embed.FS

type AnnotationView struct {
	ImageView      ImageView
	ImageInfosView ImageInfosView
	ScrollerView   ScrollerView
	image          *im.Image
	imageInfo      *a.ImageInfo
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
	v.image = &image
}

func (v *AnnotationView) DrawImageInfo(info a.ImageInfo) {
	v.imageInfo = &info
}

func (v *AnnotationView) DrawAnnotationList(annotations []an.Annotation) {
}

func (v *AnnotationView) SuccessAddBox(r addbox.Response) {
}
func (v *AnnotationView) SuccessUpdateBox(r updbox.Response) {
}
func (v *AnnotationView) SuccessDeleteAnnotation(r del.Response) {
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
				Tr(Td(v.ImageView.Render(*v.image)),
					Td(Class("align-top pl-2"), v.ImageInfosView.Render(*v.imageInfo)))),
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
