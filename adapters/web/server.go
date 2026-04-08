package web

import (
	"github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	sql_clc "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	sql_lbl "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
	list_clc "github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	list_lbl "github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
	"net/http"
)

type Server struct {
	ListCollectionsInteractor *list_clc.Interactor
	ListLabelsInteractor      *list_lbl.Interactor
	PageSize                  int
}

func (s *Server) ListCollections(w http.ResponseWriter, r *http.Request) {
	s.ListCollectionsInteractor.Execute(list_clc.Request{PageSize: s.PageSize, Page: int64(GetPageFromRequest(r))},
		NewListCollectionsPresenter(w))
}

func (s *Server) ListLabels(w http.ResponseWriter, r *http.Request) {
	s.ListLabelsInteractor.Execute(list_lbl.Request{PageSize: s.PageSize, Page: int64(GetPageFromRequest(r))},
		NewListLabelsPresenter(w))
}
func NewServer(dbPath string) *Server {
	db := sqlite.NewSQLiteDB(dbPath)
	return &Server{
		ListCollectionsInteractor: list_clc.NewInteractor(sql_clc.NewSQLiteCollectionRepo(db)),
		ListLabelsInteractor:      list_lbl.NewInteractor(sql_lbl.NewSQLiteLabelRepo(db)),
		PageSize:                  10,
	}
}
