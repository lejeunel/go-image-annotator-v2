package label

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/create"
	"net/http"
)

type CreatePresenter struct {
	Writer http.ResponseWriter
}

func (p *CreatePresenter) Success(r create.Response) {
	response := models.NewLabel{
		Name:        r.Name,
		Description: &r.Description,
	}

	json.WriteJSON(p.Writer, 200, response)

}

func (p *CreatePresenter) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}

func (p *CreatePresenter) ErrDuplication(err error) {
	json.WriteError(p.Writer, http.StatusConflict, err.Error())
}

func (p *CreatePresenter) ErrValidation(err error) {
	json.WriteError(p.Writer, http.StatusBadRequest, err.Error())
}
