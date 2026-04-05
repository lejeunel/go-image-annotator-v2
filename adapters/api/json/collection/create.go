package collection

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/create"
	"net/http"
)

type Create struct {
	Writer http.ResponseWriter
}

func (p *Create) Success(r create.Response) {
	response := models.NewCollection{
		Name:        r.Name,
		Description: &r.Description,
	}

	json.WriteJSON(p.Writer, 200, response)

}

func (p *Create) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}

func (p *Create) ErrDuplication(err error) {
	json.WriteError(p.Writer, http.StatusConflict, err.Error())
}

func (p *Create) ErrValidation(err error) {
	json.WriteError(p.Writer, http.StatusBadRequest, err.Error())
}
