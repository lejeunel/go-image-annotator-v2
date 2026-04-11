package server

import (
	u "github.com/lejeunel/go-image-annotator-v2/use-cases"
)

type Server struct {
	*u.Interactors
}

func NewServer(interactors *u.Interactors) *Server {
	return &Server{interactors}
}
