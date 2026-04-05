package label

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
	"net/http"
)

type List struct {
	Writer http.ResponseWriter
}

func (p *List) Success(r list.Response) {
	data := []models.Label{}
	for _, label := range r.Labels {
		data = append(data,
			models.Label{
				Name:        &label.Name,
				Description: &label.Description,
			})
	}

	response := models.ListLabelsResponse{Data: &data,
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
