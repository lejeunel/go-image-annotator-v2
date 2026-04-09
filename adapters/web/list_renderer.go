package web

import (
	"io"

	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
	"net/url"
)

type ListRenderer struct {
	Title       string
	NavBarSpecs html.NavBarActivatedItems
	ListURL     url.URL
	Writer      io.Writer
}

func (p ListRenderer) RenderSuccess(table html.MyTable, pagination pagination.Pagination) {
	html.MakePaginatedView(p.ListURL, p.Title, pagination, table,
		p.NavBarSpecs).Render(p.Writer)
}

func (p ListRenderer) Error(err error) {
	html.MakeErrorPage(err.Error()).Render(p.Writer)
}

func NewListRenderer(title string, listURL url.URL, navBarSpecs html.NavBarActivatedItems, w io.Writer) ListRenderer {
	return ListRenderer{Title: title, ListURL: listURL,
		NavBarSpecs: navBarSpecs, Writer: w}

}
