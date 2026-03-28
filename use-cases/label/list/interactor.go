package list

import (
	"github.com/lejeunel/go-image-annotator-v2/pagination"
)

func (i *Interactor) Execute(r Request, out OutputPort) {
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

	response := ListResponse{
		Pagination: pagination.New(int64(r.Page), r.PageSize, count),
	}
	for _, f := range found {
		response.Labels = append(response.Labels, LabelResponse{Name: f.Name, Description: f.Description})
	}
	out.Success(response)

}

type Interactor struct {
	repo Repo
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{repo: r}
}
