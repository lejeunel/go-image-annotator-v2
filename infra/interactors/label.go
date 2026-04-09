package interactors

import (
	i "github.com/lejeunel/go-image-annotator-v2/application/interactors"
	infra "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
	"github.com/lejeunel/go-image-annotator-v2/shared/validation"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/create"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/delete"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/read"
)

func NewSQLiteLabelInteractors(repo *infra.SQLiteLabelRepo) *i.LabelInteractors {
	return &i.LabelInteractors{
		Find:            *read.NewInteractor(repo),
		Create:          *create.NewInteractor(repo, validation.NewNameValidator()),
		Delete:          *delete.NewInteractor(repo),
		List:            *list.NewInteractor(repo),
		DefaultPageSize: 20,
	}
}
