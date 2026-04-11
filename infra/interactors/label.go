package interactors

import (
	infra "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
	"github.com/lejeunel/go-image-annotator-v2/shared/validation"
	lbl "github.com/lejeunel/go-image-annotator-v2/use-cases/label"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/create"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/delete"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/read"
)

func NewSQLiteLabelInteractors(repo *infra.SQLiteLabelRepo) *lbl.Interactors {
	return &lbl.Interactors{
		Find:            *read.NewInteractor(repo),
		Create:          *create.NewInteractor(repo, validation.NewNameValidator()),
		Delete:          *delete.NewInteractor(repo),
		List:            *list.NewInteractor(repo),
		DefaultPageSize: 20,
	}
}
