package web

import (
	app "github.com/lejeunel/go-image-annotator-v2/application"
	list_clc "github.com/lejeunel/go-image-annotator-v2/use-cases/collection/list"
	list_im "github.com/lejeunel/go-image-annotator-v2/use-cases/image/list"
	list_lbl "github.com/lejeunel/go-image-annotator-v2/use-cases/label/list"
)

type Server struct {
	ListCollectionsInteractor *list_clc.Interactor
	ListLabelsInteractor      *list_lbl.Interactor
	ListImagesInteractor      *list_im.Interactor
	PageSize                  int
}

func NewSQLiteServer(app *app.SQLiteApp) *Server {
	return &Server{
		ListCollectionsInteractor: list_clc.NewInteractor(app.CollectionRepo),
		ListLabelsInteractor:      list_lbl.NewInteractor(app.LabelRepo),
		ListImagesInteractor:      list_im.NewInteractor(app.ImageRepo, app.ImageStore),
		PageSize:                  10,
	}
}
