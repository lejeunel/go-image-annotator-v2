package update

type UpdateRequest struct {
	Name           string
	NewName        string
	NewDescription string
}

type UpdateResponse struct {
	Name        string
	Description string
}

type UpdateModel struct {
	Name           string
	NewName        string
	NewDescription string
}
