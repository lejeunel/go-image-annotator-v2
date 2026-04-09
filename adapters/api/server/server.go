package server

import (
	app "github.com/lejeunel/go-image-annotator-v2/application"
)

type Server struct {
	Label      *LabelServer
	Collection *CollectionServer
	Image      *ImageServer
}

func NewSQLiteServer(app *app.SQLiteApp, allowedImageFormats []string) *Server {
	return &Server{
		Label:      NewHTTPLabelServer(app.LabelRepo),
		Collection: NewHTTPCollectionServer(app.CollectionRepo),
		Image:      NewHTTPImageServer(app, allowedImageFormats),
	}
}
