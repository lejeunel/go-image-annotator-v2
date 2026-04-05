package image

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"net/http"
)

type ReadMeta struct {
	Writer http.ResponseWriter
}

func (p *ReadMeta) Success(r im.Response) {
	response := BuildImageResponse(r)
	json.WriteJSON(p.Writer, 200, response)

}

func (p *ReadMeta) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}

func (p *ReadMeta) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusBadRequest, err.Error())
}
