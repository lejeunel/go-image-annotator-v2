package image

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/list"
	"net/http"
)

type List struct {
	Writer http.ResponseWriter
}

func (p *List) Success(r list.Response) {
	response := models.ListImagesResponse{
		Pagination: models.Pagination{
			Page:       r.Pagination.Page,
			PageSize:   r.Pagination.PageSize,
			TotalItems: r.Pagination.TotalRecords,
			TotalPages: r.Pagination.TotalPages,
		},
	}

	responseImages := []models.Image{}

	for _, image := range r.Images {
		responseImages = append(responseImages, BuildImageResponse(image))
	}

	response.Images = responseImages

	json.WriteJSON(p.Writer, 200, response)

}

func (p *List) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusNotFound, err.Error())
}

func (p *List) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}
