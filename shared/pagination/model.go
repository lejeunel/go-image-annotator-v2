package pagination

import (
	"math"
)

type Pagination struct {
	TotalRecords int64
	Page         int64
	PageSize     int
	TotalPages   int64
}

func New(page int64, pageSize int, totalRecords int64) Pagination {
	return Pagination{TotalRecords: totalRecords,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int64(math.Ceil(float64(totalRecords) / float64(pageSize)))}
}
