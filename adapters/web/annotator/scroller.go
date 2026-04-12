package annotator

import (
	"fmt"
	scr "github.com/lejeunel/go-image-annotator-v2/application/scroller"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	myhtml "github.com/lejeunel/go-image-annotator-v2/shared/html"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type ScrollerView struct {
}

func MakeLink(image im.BaseImage) string {
	return fmt.Sprintf("image?id=%v&collection=%v",
		image.ImageId, image.Collection)

}

func (p *ScrollerView) Render(s scr.ScrollerState) Node {
	prevButton := myhtml.MakePreviousButton("#", false)
	nextButton := myhtml.MakeNextButton("#", false)
	if s.Previous != nil {
		prevButton = myhtml.MakePreviousButton(MakeLink(*s.Previous), true)
	}
	if s.Next != nil {
		nextButton = myhtml.MakeNextButton(MakeLink(*s.Next), true)
	}
	return Table(Tr(
		Td(prevButton),
		Td(nextButton),
	))

}
