package html

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
func MakeTitledPage(title string, content Node, scripts Node, navBarActivatedItems NavBarActivatedItems) Node {
	wrappedContent := Div(H1(Class("text-2xl font-bold text-gray-900 dark:text-gray-100"), Text(title)),
		content,
	)
	return MakeBasePage(title, wrappedContent, scripts, navBarActivatedItems)
}

func MakeBasePage(title string, content Node, scripts Node, navBarActivatedItems NavBarActivatedItems) Node {
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
		),
		Body(
			Class("bg-white text-gray-900 dark:bg-gray-900 dark:text-white"),
			MakeNavBar(navBarActivatedItems),
			Div(Class("grow w-full px-1 md:px-2 lg:px-4 py-10 md:py-20"), content),
			scripts,
		),
	))

}
