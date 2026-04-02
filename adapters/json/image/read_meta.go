package image

import (
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/image/read-meta"
	"net/http"
)

type ReadMeta struct {
	Writer http.ResponseWriter
}

func (p *ReadMeta) Success(r read_meta.Response) {
	response := models.GetImage{
		Id:         r.Id.String(),
		Collection: r.Collection,
	}

	if r.Labels != nil {
		imageLabels := []string{}
		for _, imageLabel := range r.Labels {
			imageLabels = append(imageLabels, imageLabel.Label.Name)
		}
		response.Labels = &imageLabels
	}

	if r.BoundingBoxes != nil {
		bboxes := []models.BoundingBox{}
		for _, box := range r.BoundingBoxes {
			bboxes = append(bboxes,
				models.BoundingBox{Id: box.Id.String(),
					Xc: box.Xc, Yc: box.Yc, Height: box.Height, Width: box.Width, Label: box.Label.Name})
		}
		response.BoundingBoxes = &bboxes
	}

	json.WriteJSON(p.Writer, 200, response)

}

func (p *ReadMeta) ErrInternal(err error) {
	json.WriteError(p.Writer, http.StatusInternalServerError, err.Error())
}

func (p *ReadMeta) ErrNotFound(err error) {
	json.WriteError(p.Writer, http.StatusBadRequest, err.Error())
}
