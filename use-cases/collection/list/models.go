package list

import (
	"github.com/lejeunel/go-image-annotator-v2/pagination"
)

type Request struct {
	PageSize int
	Page     int64
}

type CollectionResponse struct {
	Name        string
	Description string
}

type Response struct {
	Collections []CollectionResponse
	Pagination  pagination.Pagination
}
