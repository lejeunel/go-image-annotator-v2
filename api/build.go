package api

import (
	_ "embed"
	"github.com/lejeunel/go-image-annotator-v2/api/server"
	"html/template"
	"net/http"
)

//go:embed openapi.yaml
var openapiyaml []byte

func RegisterAPI(mux *http.ServeMux) {
	apiServer := server.NewServer()
	server.HandlerFromMuxWithBaseURL(apiServer, mux, "/api")
	specURL := "/api/openapi.yaml"
	mux.HandleFunc(specURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(openapiyaml)
	})
	mux.Handle("/api/docs", APIDocsHandler(specURL))
}

type docsData struct {
	SpecURL string
}

func APIDocsHandler(specURL string) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/docs.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, docsData{SpecURL: specURL}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
