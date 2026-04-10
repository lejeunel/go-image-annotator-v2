package html

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func BaseLibs() []Node {
	return []Node{Script(
		Src("/static/alpine.js"),
		Defer(),
	)}
}

func APIDocsLib() Node {
	return Script(Src("https://unpkg.com/@stoplight/elements/web-components.min.js"))

}

func AnnotoriousLib() []Node {
	var scripts []Node
	scripts = append(scripts, Script(Defer(), Src("/static/annotorious.js")))
	scripts = append(scripts, Link(Href("/static/annotorious.css"), Rel("stylesheet")))
	return scripts
}
