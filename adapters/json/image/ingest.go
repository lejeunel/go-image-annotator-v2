package Image

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/ingest"
	"net/http"
)

type Ingest struct {
	Writer http.ResponseWriter
}

func (p *Ingest) Success(r ingest.Response) {
	id := r.ImageId.String()
	response := models.ImageIngestionResponse{
		Id: &id,
	}

	json.WriteJSON(p.Writer, 200, response)

}

func (p *Ingest) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}

func (p *Ingest) ErrDuplication(err error) {
	json.WriteError(p.Writer, http.StatusConflict, err.Error())
}

func (p *Ingest) ErrValidation(err error) {
	json.WriteError(p.Writer, http.StatusBadRequest, err.Error())
}

func (p *Ingest) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusBadRequest, err.Error())
}
