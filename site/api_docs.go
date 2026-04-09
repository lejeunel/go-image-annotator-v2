package site

import (
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func APIDocsPage(specsFilePath string) Node {
	return html.MakeBasePage("Image Annotator API Documentation",
		Div(Class("spotlight "),
			El("elements-api",
				Attr("apiDescriptionUrl", specsFilePath),
				Attr("router", "hash"),
				Attr("layout", "sidebar"),
			)),
		html.Scripts(html.ScriptIncludes{SpotLight: true}),
		html.NavBarActivatedItems{API: true})
}
