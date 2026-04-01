package collection

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/read"
	"net/http"
)

type Find struct {
	Writer http.ResponseWriter
}

func (p *Find) Success(r read.Response) {
	response := models.Collection{
		Name:        &r.Name,
		Description: &r.Description,
	}

	json.WriteJSON(p.Writer, 200, response)

}

func (p *Find) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusNotFound, err.Error())
}

func (p *Find) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}
