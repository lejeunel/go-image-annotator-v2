package list

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
}
