package site

import (
	web "github.com/lejeunel/go-image-annotator-v2/adapters/web"
	"net/http"
)

func RegisterWebPages(mux *http.ServeMux, server web.Server) {
	mux.HandleFunc("/", HomePageHandler)
	mux.HandleFunc("/collections", server.ListCollections)
	mux.HandleFunc("/labels", server.ListLabels)
}
