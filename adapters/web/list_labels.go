package web

import (
	"net/http"

	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
	. "maragu.dev/gomponents"
)

type ListLabelsPresenter struct {
	ListRenderer
}

func (p ListLabelsPresenter) Success(r list.Response) {
	table := html.MyTable{Fields: []string{"name", "description"}}
	for _, l := range r.Labels {
		table.Rows = append(table.Rows,
			html.TableRow{Values: []Node{html.MakeTextLink("/label/"+l.Name, l.Name), Raw(l.Description)}})
	}
	p.RenderSuccess(table, r.Pagination)
}

func NewListLabelsPresenter(w http.ResponseWriter) ListLabelsPresenter {
	return ListLabelsPresenter{
		ListRenderer: NewListRenderer("Labels", "/labels",
			html.NavBarActivatedItems{Labels: true}, w),
	}
}
