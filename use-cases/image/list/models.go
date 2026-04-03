package list

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
)

type Request struct {
	CollectionName *string
	Page           int64
	PageSize       int
}

type Response struct {
	Images     []im.Response
	Pagination pagination.Pagination
}
