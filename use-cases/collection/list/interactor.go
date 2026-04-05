package list

import (
	"fmt"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	"github.com/lejeunel/go-image-annotator-v2/shared/logging"
	"github.com/lejeunel/go-image-annotator-v2/shared/pagination"
	"log/slog"
)

type Interactor struct {
	repo   Repo
	logger *slog.Logger
}

func (i *Interactor) Execute(r Request, out OutputPort) {
	found, err := i.repo.List(r)
	if err != nil {
		i.handleError(err, out)
		return
	}

	count, err := i.repo.Count()
	if err != nil {
		i.handleError(err, out)
		return
	}

	response := Response{Pagination: pagination.New(int64(r.Page), r.PageSize, *count)}
	for _, f := range found {
		response.Collections = append(response.Collections, CollectionResponse{Name: f.Name, Description: f.Description})
	}
	out.Success(response)
}

func (i *Interactor) handleError(err error, out OutputPort) {
	errCtx := "listing images"
	err = fmt.Errorf("%v: %w: %w", errCtx, err, e.ErrInternal)
	i.logger.Error(errCtx, "error", err)
	out.ErrInternal(err)
}

func NewInteractor(r Repo) *Interactor {
	return &Interactor{repo: r,
		logger: logging.NewNoOpLogger(),
	}
}
