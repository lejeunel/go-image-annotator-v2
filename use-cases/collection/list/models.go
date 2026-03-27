package list

import (
	"github.com/lejeunel/go-image-annotator-v2/pagination"
)

type Request struct {
	PageSize int
	Page     int
}

type CollectionResponse struct {
	Name        string
	Description string
}

type ListResponse struct {
	Collections []CollectionResponse
	Pagination  pagination.Pagination
}
