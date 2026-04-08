package html

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
	class := "font-medium text-on-surface underline-offset-2 hover:text-primary focus:outline-hidden focus:underline dark:text-on-surface-dark dark:hover:text-primary-dark"
	if activated {
		class = "font-bold text-primary underline-offset-2 hover:text-primary focus:outline-hidden focus:underline dark:text-primary-dark dark:hover:text-primary-dark"
	}

	return A(
		Href(url),
		Aria("current", "page"),
		Span(Class(class), Text(name)),
	)

}

func DarkModeToggle() Node {
	moon := `<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.2">
				<path stroke-linecap="round" stroke-linejoin="round" d="M21.752 15.002A9.718 9.718 0 0118 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 003 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 009.002-5.998z"></path>
			</svg>`
	sun := `<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.2">
				<path stroke-linecap="round" stroke-linejoin="round" d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z"></path>
			</svg>`
	return Button(
		Attr("@click", "toggleDark()"),
		Attr("type", "button"),
		Class(`
			whitespace-nowrap hover:bg-gray-100 dark:hover:bg-gray-800 rounded-radius px-2 py-2 text-sm font-medium tracking-wide text-surface-dark
			transition hover:opacity-75 text-center focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-surface-dark
			active:opacity-100 active:outline-offset-0 disabled:opacity-75 disabled:cursor-not-allowed
			dark:text-surface dark:focus-visible:outline-surface
		`),
		Span(
			Attr("x-html", "darkMode ? `"+sun+"` : `"+moon+"`"),
			Attr(":class", "darkMode ? 'text-gray-300' : 'text-gray-700'"),
		),
	)
}
func MakeNavBar(isActivated NavBarActivatedItems) Node {
	return Nav(
		Attr("x-on:click.away", "mobileMenuIsOpen = false"),
		Class("fixed top-0 z-30 hidden h-16 w-screen items-center justify-between border-outline px-10 py-2 backdrop-blur-xl md:flex dark:border-outline-dark bg-surface-alt/75 dark:bg-surface-dark-alt/75 border-b"),
		Aria("label", "penguin ui menu"),

		// --- Brand Logo ---
		A(
			Href("#"),
			Class("text-2xl font-bold text-on-surface-strong dark:text-on-surface-dark-strong"),
			Span(
				Text("Image"),
				Span(
					Class("text-primary dark:text-primary-dark"),
					Text("Annotation"),
				),
				Text("Platform"),
			),
		),

		// --- Desktop Menu ---
		Ul(
			Class("hidden items-center gap-4 md:flex"),

			Li(
				MakeMenuItem("Collections", "/collections", isActivated.Collections),
			),
			Li(
				MakeMenuItem("Labels", "/labels", isActivated.Labels),
			),
			Li(
				MakeMenuItem("API Docs", "/api/docs", isActivated.API),
			),

			Li(
				DarkModeToggle(),
			),
		),
	)
}
