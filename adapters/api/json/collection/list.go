package collection

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	"net/http"
)

type List struct {
	Writer http.ResponseWriter
	json.ErrorPresenter
}

func (p List) Success(r list.Response) {
	data := []models.Collection{}
	for _, c := range r.Collections {
		data = append(data,
			models.Collection{
				Name:        &c.Name,
				Description: &c.Description,
			})
	}

	response := models.ListCollectionsResponse{Data: &data,
		Pagination: json.BuildPaginationResponse(r.Pagination),
	}

	json.WriteJSON(p.Writer, 200, response)

}

func NewListPresenter(w http.ResponseWriter) List {
	return List{Writer: w, ErrorPresenter: json.ErrorPresenter{Writer: w}}
}
