package list

import im "github.com/lejeunel/go-image-annotator-v2/domain/image"

type Request struct {
	CollectionName *string
	Page           int64
	PageSize       int
}

type ImageResponse struct {
	ImageId    im.ImageId
	Collection string
}

type Pagination struct {
	Page       int64
	Total      int64
	TotalPages int64
	PageSize   int
}

type Response struct {
	Images     []*ImageResponse
	Pagination Pagination
}
