package list

import (
	"errors"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo    Repo
	service im.ImageStore
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	filteringParams := &im.FilteringParams{
		Page:       r.Page,
		PageSize:   r.PageSize,
		Collection: r.CollectionName}

	baseImages, err := i.repo.List(*filteringParams)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
		default:
			out.ErrInternal(err)
		}
		return
	}

	count, err := i.repo.Count(im.CountingParams{Collection: filteringParams.Collection})
	if err != nil {
		out.ErrInternal(err)
		return
	}

	imageResponses, ok := i.buildResponses(baseImages, out)
	if !ok {
		return
	}

	out.Success(Response{Images: imageResponses,
		Pagination: Pagination{Page: r.Page, PageSize: r.PageSize, Total: *count, TotalPages: *count / int64(r.PageSize)}})

}

func (i *Interactor) buildResponses(baseImages []*im.BaseImage, out OutputPort) ([]*ImageResponse, bool) {
	images := []*ImageResponse{}
	for _, baseImage := range baseImages {
		image, err := i.service.Find(*baseImage)
		if err != nil {
			out.ErrInternal(err)
			return nil, false
		}
		images = append(images, &ImageResponse{ImageId: image.Id, Collection: image.Collection.Name})
	}
	return images, true

}

func NewInteractor(r Repo, s im.ImageStore) *Interactor {
	return &Interactor{repo: r, service: s}
}
