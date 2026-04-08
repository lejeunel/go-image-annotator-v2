package html

import (
	s "github.com/lejeunel/go-image-annotator-v2/shared"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type ScriptIncludes struct {
	SpotLight bool
}

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

func MakePaginatedContent(baseURL string, table MyTable, p pagination.Pagination) Node {
	paginator := MakePaginator(baseURL, int(p.Page), int(p.TotalPages), len(table.Rows), int(p.TotalRecords))
	return Div(Div(Class("py-2"), paginator), table.Render())

}

func MakeTitledPage(title string, content Node, scripts Node, navBarActivatedItems NavBarActivatedItems, repoURL string) Node {
	wrappedContent := Div(Span(Class("font-extrabold text-2xl font-roboto text-gray-900 dark:text-gray-100"), Text(title)),
		Span(content),
	)
	return MakeBasePage(title, wrappedContent, scripts, navBarActivatedItems, repoURL)
}

func MakePaginatedView(baseURL string, title string, pagination pagination.Pagination,
	table MyTable, navBarActives NavBarActivatedItems) Node {

	content := MakePaginatedContent(baseURL, table, pagination)
	return MakeTitledPage(title, content,
		Scripts(ScriptIncludes{}),
		navBarActives, s.RepoURL)

}

func MakeBasePage(title string, content Node, scripts Node, navBarActivatedItems NavBarActivatedItems, repoURL string) Node {
	return Doctype(HTML(
		Attr("x-data", `{
					darkMode: false,

					init() {
console.log("init")
						this.darkMode = localStorage.getItem('dark') === 'true'
						document.documentElement.classList.toggle('dark', this.darkMode)
					},
					toggleDark() {
						this.darkMode = !this.darkMode;
						localStorage.setItem('dark', this.darkMode);
						document.documentElement.classList.toggle('dark', this.darkMode)
					}}`),
		Attr("x-init", "init()"),
		Attr("x-bind:class", "{ 'dark': darkMode }"),
		Head(
			Title(title),
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Script(Raw(`
				if (localStorage.getItem('dark') === 'true') {
					document.documentElement.classList.add('dark');
				}
			`)),
			Link(
				Rel("stylesheet"),
				Href("/static/styles.css"),
			),
			Link(Rel("stylesheet"), Href("https://fonts.googleapis.com/css2?family=Roboto&display=swap")),
		),
		Body(
			Class("bg-white text-gray-900 dark:bg-gray-900 dark:text-white"),
			MakeNavBar(navBarActivatedItems, repoURL),
			Div(Class("grow w-full px-1 md:px-2 lg:px-4 py-10 md:py-20"), content),
			scripts,
		),
	))

}
