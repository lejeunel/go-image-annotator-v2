package label

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	"net/http"
)

type Delete struct {
	Writer http.ResponseWriter
}

func (p *Delete) Success() {
	p.Writer.WriteHeader(http.StatusNoContent)

}

func (p *Delete) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusNotFound, err.Error())
}

func (p *Delete) ErrDependency(err error) {
	json.WriteError(p.Writer, http.StatusFailedDependency, err.Error())
}

func (p *Delete) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}
