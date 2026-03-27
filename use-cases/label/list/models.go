package list

import (
	"github.com/lejeunel/go-image-annotator-v2/pagination"
)

type Request struct {
	Page     int
	PageSize int
}

type LabelResponse struct {
	Name        string
	Description string
}

type ListResponse struct {
	Labels     []LabelResponse
	Pagination pagination.Pagination
}
