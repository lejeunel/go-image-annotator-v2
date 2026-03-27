package list

import (
	"github.com/lejeunel/go-image-annotator-v2/pagination"
)

func (i *ListInteractor) Execute(r Request) {
	found, err := i.repo.List(r)
	if err != nil {
		i.output.ErrInternal(err)
		return
	}

	count, err := i.repo.Count()
	if err != nil {
		i.output.ErrInternal(err)
		return
	}

	response := ListResponse{Pagination: pagination.New(int64(r.Page), r.PageSize, count)}
	for _, f := range found {
		response.Collections = append(response.Collections, CollectionResponse{Name: f.Name, Description: f.Description})
	}
	i.output.Success(response)
}

type ListInteractor struct {
	repo   Repo
	output ListOutputPort
}

func NewInteractor(r Repo, o ListOutputPort) *ListInteractor {
	return &ListInteractor{repo: r, output: o}
}
