package interactors

import (
	"github.com/jonboulle/clockwork"
	i "github.com/lejeunel/go-image-annotator-v2/application/interactors"
	infra "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	"github.com/lejeunel/go-image-annotator-v2/shared/validation"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/create"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/delete"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/read"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/update"
)

func NewSQLiteCollectionInteractors(repo *infra.SQLiteCollectionRepo) *i.CollectionInteractors {
	return &i.CollectionInteractors{
		Find:            *read.NewInteractor(repo),
		Create:          *create.NewInteractor(repo, validation.NewNameValidator(), clockwork.NewRealClock()),
		Delete:          *delete.NewInteractor(repo),
		List:            *list.NewInteractor(repo),
		Update:          *update.NewInteractor(repo),
		DefaultPageSize: 20,
	}
}
