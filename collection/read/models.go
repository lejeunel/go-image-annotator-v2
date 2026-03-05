package read

type ReadCollectionRequest struct {
	Name string
}

type ReadCollectionResponse struct {
	Name        string
	Description string
}

type Collection struct {
	Name        string
	Description string
}
