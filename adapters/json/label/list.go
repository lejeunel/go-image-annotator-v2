package label

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
	"net/http"
)

type ListPresenter struct {
	Writer http.ResponseWriter
}

func (p *ListPresenter) Success(r list.Response) {
	data := []models.Label{}
	for _, label := range r.Labels {
		data = append(data,
			models.Label{
				Name:        &label.Name,
				Description: &label.Description,
			})
	}

	response := models.ListLabelsResponse{Data: &data,
		Pagination: &models.Pagination{
			Page:       &r.Pagination.Page,
			PageSize:   &r.Pagination.PageSize,
			TotalItems: &r.Pagination.TotalRecords,
			TotalPages: &r.Pagination.TotalPages,
		},
	}

	json.WriteJSON(p.Writer, 200, response)

}

func (p *ListPresenter) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusNotFound, err.Error())
}

func (p *ListPresenter) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}
