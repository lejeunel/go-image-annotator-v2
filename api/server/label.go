package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/lejeunel/go-image-annotator-v2/adapters/json"
	pres "github.com/lejeunel/go-image-annotator-v2/adapters/json/label"
	"github.com/lejeunel/go-image-annotator-v2/api/models"
	infra "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/create"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/delete"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
	"github.com/lejeunel/go-image-annotator-v2/use-cases/label/read"
	"github.com/lejeunel/go-image-annotator-v2/validation"
)

type LabelServer struct {
	Find            read.Interactor
	Create          create.Interactor
	Delete          delete.Interactor
	List            list.Interactor
	DefaultPageSize int
}

func NewHTTPLabelServer(db *sqlx.DB) *LabelServer {
	repo := infra.NewSQLiteLabelRepo(db)
	return &LabelServer{
		Find:            *read.NewInteractor(repo),
		Create:          *create.NewInteractor(repo, validation.NewNameValidator()),
		Delete:          *delete.NewInteractor(repo),
		List:            *list.NewInteractor(repo),
		DefaultPageSize: 20,
	}
}

func (s *Server) FindLabelByName(
	w http.ResponseWriter,
	r *http.Request,
	name string,
) {

	s.Label.Find.Execute(read.Request{Name: name}, &pres.FindPresenter{Writer: w})
}

func (s *Server) CreateLabel(
	w http.ResponseWriter,
	r *http.Request,
) {
	body, ok := json.DecodeJSONOrFail[models.NewLabel](w, r)
	if !ok {
		return
	}

	s.Label.Create.Execute(create.Request{Name: body.Name, Description: *body.Description},
		&pres.CreatePresenter{Writer: w})
}
func (s *Server) DeleteLabelByName(w http.ResponseWriter, r *http.Request, name string) {
	s.Label.Delete.Execute(delete.Request{Name: name}, &pres.DeletePresenter{Writer: w})

}
func (s *Server) ListLabels(w http.ResponseWriter, r *http.Request, params ListLabelsParams) {
	req := list.Request{Page: 1, PageSize: s.Label.DefaultPageSize}
	if p := params.Page; p != nil {
		req.Page = *p
	}
	if p := params.PageSize; p != nil {
		req.PageSize = *p
	}
	s.Label.List.Execute(req, &pres.ListPresenter{Writer: w})

}
