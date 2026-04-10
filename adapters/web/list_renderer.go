package web

import (
	"io"

	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	n "github.com/lejeunel/go-image-annotator-v2/shared/navigation"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
	"net/url"
)

type ListRenderer struct {
	Title      string
	ActivePage n.ActivePage
	ListURL    url.URL
	Writer     io.Writer
}

func (p ListRenderer) RenderSuccess(table html.MyTable, pagination pagination.Pagination) {

	content := html.MakePaginatedContent(p.ListURL, table, pagination)
	html.NewTitledPageBuilder(p.Title).SetContent(content).SetActive(p.ActivePage).Render(p.Writer)
}

func (p ListRenderer) Error(err error) {
	html.NewPageBuilder().SetError(err.Error()).Render(p.Writer)
}

func NewListRenderer(title string, listURL url.URL, page n.ActivePage, w io.Writer) ListRenderer {
	return ListRenderer{Title: title, ListURL: listURL,
		ActivePage: page, Writer: w}

}
