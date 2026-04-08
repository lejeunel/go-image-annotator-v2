package web

import (
	"io"

	s "github.com/lejeunel/go-image-annotator-v2/shared"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
	. "maragu.dev/gomponents"
)

type ListRenderer struct {
	Title       string
	NavBarSpecs html.NavBarActivatedItems
	ListURL     string
	Writer      io.Writer
}

func (p ListRenderer) RenderSuccess(table html.MyTable, pagination pagination.Pagination) {
	html.MakePaginatedView(p.ListURL, p.Title, pagination, table,
		p.NavBarSpecs).Render(p.Writer)
}

func (p ListRenderer) Error(err error) {
	html.MakeBasePage(p.Title, Text(err.Error()),
		html.Scripts(html.ScriptIncludes{}), p.NavBarSpecs, s.RepoURL).Render(p.Writer)
}

func NewListRenderer(title, listURL string, navBarSpecs html.NavBarActivatedItems, w io.Writer) ListRenderer {
	return ListRenderer{Title: title, ListURL: listURL,
		NavBarSpecs: navBarSpecs, Writer: w}

}
