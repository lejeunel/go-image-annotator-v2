package json

import (
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
)

func BuildPaginationResponse(p pagination.Pagination) models.Pagination {
	return models.Pagination{
		Page:       p.Page,
		PageSize:   p.PageSize,
		TotalItems: p.TotalRecords,
		TotalPages: p.TotalPages,
	}

}
