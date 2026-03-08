package read

func (i *ListInteractor) Execute(r ListRequest) {
	found, err := i.repo.List(r)
	if err != nil {
		i.output.ErrInternal(err)
		return
	}

	response := ListResponse{}
	for _, f := range found {
		response.Collections = append(response.Collections, CollectionResponse{Name: f.Name, Description: f.Description})
	}
	i.output.Success(response)

}

type ListInteractor struct {
	repo   ListRepo
	output ListOutputPort
}

func NewListInteractor(r ListRepo, o ListOutputPort) *ListInteractor {
	return &ListInteractor{repo: r, output: o}
}
