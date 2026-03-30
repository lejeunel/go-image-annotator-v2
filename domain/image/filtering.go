package image

type FilteringParams struct {
	Collection *string
	PageSize   int
	Page       int64
}

type CountingParams struct {
	Collection *string
}
