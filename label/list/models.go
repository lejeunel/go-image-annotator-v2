package read

type ListRequest struct {
	Page     int
	PageSize int
}

type LabelResponse struct {
	Name        string
	Description string
}

type ListResponse struct {
	Labels []LabelResponse
}
