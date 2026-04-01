package server

import (
	"github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
)

type Server struct {
	Label      *LabelServer
	Collection *CollectionServer
}

func NewServer(dbPath string) *Server {
	db := sqlite.NewSQLiteDB(dbPath)
	return &Server{
		Label:      NewHTTPLabelServer(db),
		Collection: NewHTTPCollectionServer(db),
	}
}
