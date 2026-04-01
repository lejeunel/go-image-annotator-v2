package server

import (
	"github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
)

type Server struct {
	Label *LabelServer
}

func NewServer() *Server {
	db := sqlite.NewSQLiteDB(":memory:")
	return &Server{Label: NewHTTPLabelServer(db)}
}
