package annotator

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	. "maragu.dev/gomponents"
)

type ImageInfosView struct {
	result Node
}

func (p *ImageInfosView) Success(image *im.Image) {
	table := html.SpecTable{}
	table.Rows = append(table.Rows, html.SpecTableRow{Name: "id", Value: image.Id.String()})
	table.Rows = append(table.Rows, html.SpecTableRow{Name: "collection", Value: image.Collection.Name})
	tableNode := table.Render()
	p.result = tableNode

}
func (p *ImageInfosView) Error(err error) {
	p.result = Text(err.Error())
}

func (p *ImageInfosView) Build() Node {
	return p.result

}
