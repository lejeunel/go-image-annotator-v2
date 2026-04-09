package web

import (
	"fmt"
	"net/http"
	"net/url"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	html "github.com/lejeunel/go-image-annotator-v2/shared/html"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/list"
	list_im "github.com/lejeunel/go-image-annotator-v2/use-cases/image/list"
	. "maragu.dev/gomponents"
)

type ListImagesPresenter struct {
	ListRenderer
}

func (p ListImagesPresenter) Success(r list.Response) {
	table := html.MyTable{Fields: []string{"id", "collection"}}
	for _, im := range r.Images {
		table.Rows = append(table.Rows,
			html.TableRow{Values: []Node{Text(im.Id.String()),
				Text(im.Collection)}})
	}
	p.RenderSuccess(table, r.Pagination)
}

func (s *Server) ListImages(w http.ResponseWriter, r *http.Request) {

	collection := r.URL.Query().Get("collection")
	if collection == "" {
		html.MakeErrorPage(fmt.Errorf("parsing url to get collection name: %w", e.ErrURLParsing).Error()).Render(w)
	}
	s.ListImagesInteractor.Execute(list_im.Request{PageSize: s.PageSize,
		Page:           int64(GetPageFromRequest(r)),
		CollectionName: &collection},
		NewListImagesPresenter(w, *r.URL))
}

func NewListImagesPresenter(w http.ResponseWriter, baseURL url.URL) ListImagesPresenter {
	return ListImagesPresenter{
		ListRenderer: NewListRenderer("Images", baseURL,
			html.NavBarActivatedItems{}, w),
	}
}
