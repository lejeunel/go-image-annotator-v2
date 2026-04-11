package web

import (
	a "github.com/lejeunel/go-image-annotator-v2/application/annotator"
	u "github.com/lejeunel/go-image-annotator-v2/use-cases"
)

type Server struct {
	*u.Interactors
	annotatorBuilder *a.AnnotatorBuilder
}

func NewServer(interactors *u.Interactors, annotatorBuilder *a.AnnotatorBuilder) *Server {
	return &Server{Interactors: interactors, annotatorBuilder: annotatorBuilder}
}
