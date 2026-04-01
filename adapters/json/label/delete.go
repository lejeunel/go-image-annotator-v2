package label

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"net/http"
)

type DeletePresenter struct {
	Writer http.ResponseWriter
}

func (p *DeletePresenter) Success() {
	p.Writer.WriteHeader(http.StatusNoContent)

}

func (p *DeletePresenter) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusNotFound, err.Error())
}

func (p *DeletePresenter) ErrDependency(err error) {
	json.WriteError(p.Writer, http.StatusFailedDependency, err.Error())
}

func (p *DeletePresenter) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}
