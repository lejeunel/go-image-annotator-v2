package annotator

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	. "maragu.dev/gomponents"
)

type ImageInfosView struct {
	result Node
}

func (p *ImageInfosView) Render(image *im.Image) Node {
	table := html.SpecTable{}
	table.Rows = append(table.Rows, html.SpecTableRow{Name: "id", Value: image.Id.String()})
	table.Rows = append(table.Rows, html.SpecTableRow{Name: "collection", Value: image.Collection.Name})
	return table.Render()

}
