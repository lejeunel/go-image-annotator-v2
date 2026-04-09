package server

import (
	i "github.com/lejeunel/go-image-annotator-v2/application/interactors"
)

type Server struct {
	*i.Interactors
}

func NewServer(interactors *i.Interactors) *Server {
	return &Server{interactors}
}
