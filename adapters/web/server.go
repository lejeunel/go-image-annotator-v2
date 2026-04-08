package web

import (
	"github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	sql_clc "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	list_clc "github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	"net/http"
)

type Server struct {
	Interactor *list_clc.Interactor
}

func (s *Server) ListCollections(w http.ResponseWriter, r *http.Request) {
	s.Interactor.Execute(list_clc.Request{PageSize: 999, Page: 1}, NewListPresenter(w))
}

func NewServer(dbPath string) *Server {
	db := sqlite.NewSQLiteDB(dbPath)
	return &Server{
		Interactor: list_clc.NewInteractor(sql_clc.NewSQLiteCollectionRepo(db)),
	}
}
