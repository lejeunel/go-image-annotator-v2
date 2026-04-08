package list

import (
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
)

type Request struct {
	PageSize int
	Page     int64
}

type Response struct {
	Collections []*clc.Collection
	Pagination  pagination.Pagination
}
