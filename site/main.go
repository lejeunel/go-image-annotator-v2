package site

import (
	api "github.com/lejeunel/go-image-annotator-v2/adapters/api/server"
	web "github.com/lejeunel/go-image-annotator-v2/adapters/web"
	"net/http"
)

type SiteConfig struct {
	APIDocsPath      string
	OpenAPISpecsPath string
}

func RegisterHandlers(mux *http.ServeMux, apiServer api.Server, webServer web.Server, cfg SiteConfig) {
	RegisterAPI(mux, apiServer, cfg.APIDocsPath, cfg.OpenAPISpecsPath)
	RegisterStaticFiles(mux)
	RegisterWebPages(mux, webServer)
}
