package image

import (
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

func BuildImageResponse(image im.ImageResponse) models.Image {
	response := models.Image{
		Id:         image.Id.String(),
		Collection: image.Collection,
	}
	labelsToAdd := []string{}
	if image.Labels != nil {
		for _, l := range image.Labels {
			labelsToAdd = append(labelsToAdd, l.Label.Name)
		}
		response.Labels = &labelsToAdd
	}

	if image.BoundingBoxes != nil {
		boxesToAdd := []models.BoundingBox{}
		for _, b := range image.BoundingBoxes {
			boxesToAdd = append(boxesToAdd,
				models.BoundingBox{Id: b.Id.String(),
					Xc: b.Xc, Yc: b.Yc, Height: b.Height, Width: b.Width, Label: b.Label.Name})
		}
		response.BoundingBoxes = &boxesToAdd
	}

	return response

}
