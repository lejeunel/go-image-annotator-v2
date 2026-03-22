package list

import (
	"errors"

	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo    Repo
	output  OutputPort
	service im.ImageService
}

func (i *Interactor) Execute(r Request) {
	filteringParams := &FilteringParams{Page: r.Page, PageSize: r.PageSize}

	if ok := i.processCollectionName(r.CollectionName, filteringParams); !ok {
		return
	}

	baseImages, err := i.repo.List(*filteringParams)
	if err != nil {
		i.output.ErrInternal(err)
		return
	}

	count, err := i.repo.Count(*filteringParams)
	if err != nil {
		i.output.ErrInternal(err)
		return
	}

	imageResponses, ok := i.buildResponses(baseImages)
	if !ok {
		return
	}

	i.output.Success(Response{Images: imageResponses,
		Pagination: Pagination{Page: r.Page, PageSize: r.PageSize, Total: *count, TotalPages: *count / r.PageSize}})

}

func (i *Interactor) buildResponses(baseImages []*im.BaseImage) ([]*ImageResponse, bool) {
	images := []*ImageResponse{}
	for _, baseImage := range baseImages {
		image, err := i.service.Find(*baseImage)
		if err != nil {
			i.output.ErrInternal(err)
			return nil, false
		}
		images = append(images, &ImageResponse{ImageId: image.Id, Collection: image.Collection.Name})
	}
	return images, true

}

func (i *Interactor) processCollectionName(collectionName *string, f *FilteringParams) bool {
	if collectionName != nil {
		collectionId, err := i.repo.FindCollectionIdByName(*collectionName)
		if err != nil {
			switch {
			case errors.Is(err, e.ErrNotFound):
				i.output.ErrNotFound(err)
			default:
				i.output.ErrInternal(err)
			}
			return false
		}
		f.CollectionId = collectionId
	}
	return true

}

func NewInteractor(r Repo, o OutputPort, s im.ImageService) *Interactor {
	return &Interactor{repo: r, output: o, service: s}
}
