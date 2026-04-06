package site

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type NavBarActivatedItems struct {
	Collections bool
	Labels      bool
	API         bool
}

func MakeMenuItem(name string, url string, activated bool) Node {
	class := "font-medium text-primary underline-offset-2 hover:text-primary focus:outline-hidden focus:underline dark:text-primary-dark dark:hover:text-primary-dark"
	if activated {
		class = "font-bold text-primary underline-offset-2 hover:text-primary focus:outline-hidden focus:underline dark:text-primary-dark dark:hover:text-primary-dark"
	}
	return A(
		Href(url),
		Class(class),
		Class("font-bold text-primary underline-offset-2 hover:text-primary focus:outline-hidden focus:underline dark:text-primary-dark dark:hover:text-primary-dark"),
		Aria("current", "page"),
		Text(name),
	)

}

func MakeNavBar(isActivated NavBarActivatedItems) Node {
	return Nav(
		Attr("x-data", "{ mobileMenuIsOpen: false }"),
		Attr("x-on:click.away", "mobileMenuIsOpen = false"),
		Class("flex items-center justify-between border-b border-outline px-6 py-4 dark:border-outline-dark"),
		Aria("label", "penguin ui menu"),

		// --- Brand Logo ---
		A(
			Href("#"),
			Class("text-2xl font-bold text-on-surface-strong dark:text-on-surface-dark-strong"),
			Span(
				Text("Image"),
				Span(
					Class("text-primary text-blue-600 dark:text-blue-600"),
					Text("Annotation"),
				),
				Text("Platform"),
			),
		),

		// --- Desktop Menu ---
		Ul(
			Class("hidden items-center gap-4 md:flex"),

			Li(
				MakeMenuItem("Collections", "collections", isActivated.Collections),
			),
			Li(
				MakeMenuItem("Labels", "labels", isActivated.Labels),
			),
			Li(
				MakeMenuItem("API Docs", "/api/docs", isActivated.API),
			),
		),
	)
}
