package server

import (
	"github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
)

type Server struct {
	Label      *LabelServer
	Collection *CollectionServer
	Image      *ImageServer
}

func NewServer(dbPath, artefactDir string, allowedImageFormats []string) *Server {
	db := sqlite.NewSQLiteDB(dbPath)
	return &Server{
		Label:      NewHTTPLabelServer(db),
		Collection: NewHTTPCollectionServer(db),
		Image:      NewHTTPImageServer(db, artefactDir, allowedImageFormats),
	}
}
