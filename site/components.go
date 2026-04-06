package site

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type ScriptIncludes struct {
	SpotLight bool
}

// Scripts returns a group of <script> tags.
// If includeAlpine is true, Alpine.js is added.
func Scripts(include ScriptIncludes) Node {
	var scripts []Node

	if include.SpotLight {
		scripts = append(scripts,
			Script(Src("https://unpkg.com/@stoplight/elements/web-components.min.js")),
		)

	}

	scripts = append(scripts,
		Script(
			Src("https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"),
			Defer(),
		),
	)

	return Group(scripts)
}

func MakePage(title string, content Node, additionalCSS Node, scripts Node, navBarActivatedItems NavBarActivatedItems) Node {
	return HTML(
		Head(
			Title(title),
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Link(
				Rel("stylesheet"),
				Href("/static/styles.css"),
			),
			additionalCSS,
		),
		Body(
			MakeNavBar(navBarActivatedItems),
			content,
			scripts,
		),
	)

}
