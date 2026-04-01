package label

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/read"
	"net/http"
)

type FindPresenter struct {
	Writer http.ResponseWriter
}

func (p *FindPresenter) Success(r read.Response) {
	response := models.Label{
		Name:        &r.Name,
		Description: &r.Description,
	}

	json.WriteJSON(p.Writer, 200, response)

}

func (p *FindPresenter) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusNotFound, err.Error())
}

func (p *FindPresenter) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}
