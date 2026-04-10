package site

import (
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	n "github.com/lejeunel/go-image-annotator-v2/shared/navigation"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func APIDocsPage(specsFilePath string) Node {
	p := html.NewPageBuilder()
	p.AddScripts(html.APIDocsLib())
	p.SetContent(Div(Class("spotlight "),
		El("elements-api",
			Attr("apiDescriptionUrl", specsFilePath),
			Attr("router", "hash"),
			Attr("layout", "sidebar"),
		)))
	p.SetActive(n.APIDocsPageActive)
	return p.Build()

}
