package site

import (
	"fmt"
	api "github.com/lejeunel/go-image-annotator-v2/adapters/api/server"
	web "github.com/lejeunel/go-image-annotator-v2/adapters/web"
	"github.com/lejeunel/go-image-annotator-v2/config"
	"github.com/lejeunel/go-image-annotator-v2/infra"
	i "github.com/lejeunel/go-image-annotator-v2/infra/interactors"
	"net/http"
)

type SiteConfig struct {
	APIDocsPath      string
	OpenAPISpecsPath string
}

func RegisterHandlers(mux *http.ServeMux, apiServer api.Server, webServer web.Server, cfg SiteConfig) {
	RegisterAPI(mux, apiServer, cfg.APIDocsPath, cfg.OpenAPISpecsPath)
	RegisterStaticFiles(mux)
	web.RegisterWebPages(mux, webServer)
}

func Serve(port int) {
	cfg := config.Parse()
	mux := http.NewServeMux()

	interactors := i.NewSQLiteInteractors(infra.NewSQLiteInfra(cfg.DBPath, cfg.ArtefactDir),
		cfg.AllowedImageFormats)
	RegisterHandlers(mux,
		*api.NewServer(interactors),
		*web.NewServer(interactors),
		SiteConfig{APIDocsPath: "/api/docs", OpenAPISpecsPath: "/api/openapi.yaml"})

	fmt.Println("serving on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
