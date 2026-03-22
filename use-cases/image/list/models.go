package list

import im "github.com/lejeunel/go-image-annotator-v2/domain/image"

type Request struct {
	CollectionName *string
	Page           int
	PageSize       int
}

type ImageResponse struct {
	ImageId    im.ImageId
	Collection string
}

type Pagination struct {
	Page       int
	Total      int
	TotalPages int
	PageSize   int
}

type Response struct {
	Images     []*ImageResponse
	Pagination Pagination
}
