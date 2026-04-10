package html

import (
	"net/url"

	n "github.com/lejeunel/go-image-annotator-v2/shared/navigation"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type ScriptIncludes struct {
	SpotLight   bool
	Annotorious bool
}

func DefaultScripts() []Node {
	return []Node{Script(
		Src("/static/alpine.js"),
		Defer(),
	)}
}

func APIDocsScripts() Node {
	return Script(Src("https://unpkg.com/@stoplight/elements/web-components.min.js"))

}

func AnnotoriousScripts() []Node {
	var scripts []Node
	scripts = append(scripts, Script(Defer(), Src("/static/annotorious.js")))
	scripts = append(scripts, Link(Href("/static/annotorious.css"), Rel("stylesheet")))
	return scripts
}

func Scripts(include ScriptIncludes) Node {
	var scripts []Node

	if include.SpotLight {
		scripts = append(scripts, APIDocsScripts())
	}

	scripts = append(scripts, DefaultScripts()...)

	if include.Annotorious {
		scripts = append(scripts, Script(Defer(), Src("/static/annotorious.js")))
		scripts = append(scripts, Link(Href("/static/annotorious.css"), Rel("stylesheet")))
	}
	return Group(scripts)
}

func MakePaginatedContent(baseURL url.URL, table MyTable, p pagination.Pagination) Node {
	paginator := MakePaginator(baseURL, int(p.Page), int(p.TotalPages), len(table.Rows), int(p.TotalRecords))
	return Div(Div(Class("py-2"), paginator), table.Render())

}

func MakePaginatedView(baseURL url.URL, title string, pagination pagination.Pagination,
	table MyTable, activePage n.ActivePage) Node {

	content := MakePaginatedContent(baseURL, table, pagination)
	p := NewTitledPageBuilder(title)
	p.SetContent(content)
	p.SetActive(activePage)
	return p.Build()
}
