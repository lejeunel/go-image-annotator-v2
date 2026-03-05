package update

type UpdateCollectionRequest struct {
	Name           string
	NewName        string
	NewDescription string
}

type UpdateCollectionResponse struct {
	Name        string
	Description string
}
