package site

import (
	"fmt"
	api "github.com/lejeunel/go-image-annotator-v2/adapters/api/server"
	web "github.com/lejeunel/go-image-annotator-v2/adapters/web"
	app "github.com/lejeunel/go-image-annotator-v2/application"
	"github.com/lejeunel/go-image-annotator-v2/config"
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

	app := app.NewSQLiteApp(cfg.DBPath, cfg.ArtefactDir)
	RegisterHandlers(mux,
		*api.NewSQLiteServer(app, cfg.AllowedImageFormats),
		*web.NewSQLiteServer(app),
		SiteConfig{APIDocsPath: "/api/docs", OpenAPISpecsPath: "/api/openapi.yaml"})

	fmt.Println("serving on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
