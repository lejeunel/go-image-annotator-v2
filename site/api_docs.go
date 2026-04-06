package site

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// APIDocsPage renders the Stoplight Elements API documentation page.
func APIDocsPage(specsFilePath string) Node {
	return MakePage("Image Annotator API Documentation",
		El("elements-api",
			Attr("apiDescriptionUrl", specsFilePath),
			Attr("router", "hash"),
			Attr("layout", "sidebar"),
		), Link(
			Rel("stylesheet"),
			Href("https://unpkg.com/@stoplight/elements/styles.min.css"),
		),
		Scripts(ScriptIncludes{SpotLight: true}),
		NavBarActivatedItems{API: true},
	)
}
