package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/json"
	presenter "github.com/lejeunel/go-image-annotator-v2/adapters/api/json/collection"
	"github.com/lejeunel/go-image-annotator-v2/adapters/api/models"
	infra "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	"github.com/lejeunel/go-image-annotator-v2/shared/validation"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/create"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/delete"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/collection/read"
)

type CollectionServer struct {
	Find            read.Interactor
	Create          create.Interactor
	Delete          delete.Interactor
	List            list.Interactor
	DefaultPageSize int
}

func NewHTTPCollectionServer(db *sqlx.DB) *CollectionServer {
	repo := infra.NewSQLiteCollectionRepo(db)
	return &CollectionServer{
		Find:            *read.NewInteractor(repo),
		Create:          *create.NewInteractor(repo, validation.NewNameValidator()),
		Delete:          *delete.NewInteractor(repo),
		List:            *list.NewInteractor(repo),
		DefaultPageSize: 20,
	}
}

func (s *Server) FindCollectionByName(w http.ResponseWriter, r *http.Request, name string) {

	s.Collection.Find.Execute(read.Request{Name: name}, presenter.NewFindPresenter(w))
}
func (s *Server) CreateCollection(w http.ResponseWriter, r *http.Request) {
	body, ok := json.DecodeJSONOrFail[models.NewCollection](w, r)
	if !ok {
		return
	}

	s.Collection.Create.Execute(create.Request{Name: body.Name, Description: *body.Description},
		presenter.NewCreatePresenter(w))
}
func (s *Server) DeleteCollectionByName(w http.ResponseWriter, r *http.Request, name string) {
	s.Collection.Delete.Execute(delete.Request{Name: name}, presenter.NewDeletePresenter(w))

}
func (s *Server) ListCollections(w http.ResponseWriter, r *http.Request, params ListCollectionsParams) {
	req := list.Request{Page: 1, PageSize: s.Collection.DefaultPageSize}
	if p := params.Page; p != nil {
		req.Page = *p
	}
	if p := params.PageSize; p != nil {
		req.PageSize = *p
	}
	s.Collection.List.Execute(req, presenter.NewListPresenter(w))

}
