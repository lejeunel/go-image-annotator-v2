package web

import (
	"net/http"

	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	. "maragu.dev/gomponents"
)

type ListCollectionsPresenter struct {
	ListRenderer
}

func (p ListCollectionsPresenter) Success(r list.Response) {
	table := html.MyTable{Fields: []string{"name", "description", "created"}}
	for _, c := range r.Collections {
		table.Rows = append(table.Rows,
			html.TableRow{Values: []Node{html.MakeTextLink("/collection/"+c.Name, c.Name),
				Raw(c.Description), Raw(DateTimeToStr(c.CreatedAt))}})
	}
	p.RenderSuccess(table, r.Pagination)
}

func NewListCollectionsPresenter(w http.ResponseWriter) ListCollectionsPresenter {
	return ListCollectionsPresenter{
		ListRenderer: NewListRenderer("Collections", "/collections",
			html.NavBarActivatedItems{Collections: true}, w),
	}
}
