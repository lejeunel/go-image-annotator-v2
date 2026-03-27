package list

import (
	"github.com/lejeunel/go-image-annotator-v2/pagination"
)

func (i *Interactor) Execute(r Request) {
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

	response := ListResponse{
		Pagination: pagination.New(int64(r.Page), r.PageSize, count),
	}
	for _, f := range found {
		response.Labels = append(response.Labels, LabelResponse{Name: f.Name, Description: f.Description})
	}
	i.output.Success(response)

}

type Interactor struct {
	repo   Repo
	output OutputPort
}

func NewInteractor(r Repo, o OutputPort) *Interactor {
	return &Interactor{repo: r, output: o}
}
