package web

import (
	"net/http"

	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	. "maragu.dev/gomponents"
)

type ListPresenter struct {
	Writer http.ResponseWriter
	Title  string
}

func (p ListPresenter) Success(r list.Response) {
	var rows []html.TableRow
	for _, c := range r.Collections {
		rows = append(rows,
			html.TableRow{Values: []Node{Raw(c.Name), Raw(c.Description)}})
	}

	html.MakeTitledPage(p.Title, html.MyTable([]string{"name", "description"}, rows),
		html.Scripts(html.ScriptIncludes{}),
		html.NavBarActivatedItems{Collections: true}).Render(p.Writer)
}

func (p ListPresenter) Error(err error) {
	html.MakeBasePage(p.Title, Text(err.Error()), html.Scripts(html.ScriptIncludes{}), html.NavBarActivatedItems{})
}

func NewListPresenter(w http.ResponseWriter) ListPresenter {
	return ListPresenter{Writer: w, Title: "Collections"}
}
