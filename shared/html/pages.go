package html

import (
	"net/url"

	n "github.com/lejeunel/go-image-annotator-v2/shared/navigation"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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
