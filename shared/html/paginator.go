package html

import (
	"fmt"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"strconv"
)

func MakePaginatorEllipsis() Node {
	return Li(A(Href("#"),
		Class("flex size-6 items-center justify-center rounded-radius p-1 text-on-surface hover:text-primary dark:text-on-surface-dark dark:hover:text-primary-dark"),
		Aria("label", "more-pages"),
		Raw(`
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" aria-hidden="true" stroke="currentColor" class="size-6">
					<path stroke-linecap="round" stroke-linejoin="round" d="M6.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM18.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
				</svg>`),
	))
}

func MakePaginatorNumberedButton(baseURL string, pageNumber int, isActive bool) Node {
	url := fmt.Sprintf("%v?page=%v", baseURL, pageNumber)
	class := Class("flex size-6 items-center justify-center rounded-radius p-1 text-on-surface hover:text-primary dark:text-on-surface-dark dark:hover:text-primary-dark")
	if isActive {
		class = Class("flex size-6 items-center justify-center rounded-radius bg-primary p-1 font-bold text-on-primary dark:bg-primary-dark dark:text-on-primary-dark")
	}
	return Li(A(Href(url),
		class,
		Aria("label", fmt.Sprintf("page %v", pageNumber)),
		Text(strconv.Itoa(pageNumber)),
	))
}

func MakePaginatorPrevious(url string) Node {
	return Li(A(Href(url), Class("flex items-center rounded-radius p-1 text-on-surface hover:text-primary dark:text-on-surface-dark dark:hover:text-primary-dark"),
		Raw(`
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="size-6">
						<path fill-rule="evenodd" d="M11.78 5.22a.75.75 0 0 1 0 1.06L8.06 10l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0Z" clip-rule="evenodd" />
					</svg>

				`),
		Text("Previous"),
	))
}

func MakePaginatorNext(url string) Node {
	return Li(A(Href(url), Class("flex items-center rounded-radius p-1 text-on-surface hover:text-primary dark:text-on-surface-dark dark:hover:text-primary-dark"),
		Text("Next"),
		Raw(`
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="size-6">
						<path fill-rule="evenodd" d="M8.22 5.22a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.94 10 8.22 6.28a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd" />
					</svg>

				`),
	))

}

func MakePaginator(baseURL string, currentPage, lastPage, numItems, totalItems int) Node {
	prevURL := fmt.Sprintf("%v?page=%v", baseURL, currentPage-1)
	nextURL := fmt.Sprintf("%v?page=%v", baseURL, currentPage+1)
	return Nav(Aria("label", "pagination"),
		Ul(Class("flex shrink-0 items-center gap-2 text-sm font-medium"),
			If(currentPage > 1, MakePaginatorPrevious(prevURL)),
			If(currentPage > 1, MakePaginatorNumberedButton(baseURL, currentPage-1, false)),
			If(lastPage > 1, MakePaginatorNumberedButton(baseURL, currentPage, true)),
			If(lastPage-2 > currentPage, MakePaginatorEllipsis()),
			If(lastPage-1 > currentPage, MakePaginatorNumberedButton(baseURL, lastPage-1, false)),
			If(lastPage > currentPage, MakePaginatorNumberedButton(baseURL, lastPage, false)),
			If(currentPage < lastPage, MakePaginatorNext(nextURL)),
			Span(Class("font-light"), Text(fmt.Sprintf("Showing %v items out of %v", numItems, totalItems))),
		))
}
