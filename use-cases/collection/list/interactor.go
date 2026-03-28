package list

import (
	"github.com/lejeunel/go-image-annotator-v2/pagination"
)

func (i *ListInteractor) Execute(r Request, out OutputPort) {
	found, err := i.repo.List(r)
	if err != nil {
		out.ErrInternal(err)
		return
	}

	count, err := i.repo.Count()
	if err != nil {
		out.ErrInternal(err)
		return
	}

	response := ListResponse{Pagination: pagination.New(int64(r.Page), r.PageSize, count)}
	for _, f := range found {
		response.Collections = append(response.Collections, CollectionResponse{Name: f.Name, Description: f.Description})
	}
	out.Success(response)
}

type ListInteractor struct {
	repo Repo
}

func NewInteractor(r Repo) *ListInteractor {
	return &ListInteractor{repo: r}
}
