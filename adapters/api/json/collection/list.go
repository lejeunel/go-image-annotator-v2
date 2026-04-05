package collection

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	"net/http"
)

type List struct {
	Writer http.ResponseWriter
}

func (p *List) Success(r list.Response) {
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

func (p *List) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusNotFound, err.Error())
}

func (p *List) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}
