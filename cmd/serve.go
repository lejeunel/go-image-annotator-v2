package cmd

import (
	"fmt"
	apiServer "github.com/lejeunel/go-image-annotator-v2/adapters/api/server"
	"github.com/lejeunel/go-image-annotator-v2/config"
	"github.com/lejeunel/go-image-annotator-v2/site"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	port     int
	ServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "Run server",
		Run: func(cmd *cobra.Command, args []string) {
			serve(port)
		},
	}
)

func init() {
	ServeCmd.Flags().IntVarP(&port, "port", "p", 80, "port to serve on")
}

func serve(port int) {
	cfg := config.Parse()
	mux := http.NewServeMux()

	siteConfig := site.SiteConfig{APIDocsPath: "/api/docs", OpenAPISpecsPath: "/api/openapi.yaml"}

	site.RegisterHandlers(mux,
		*apiServer.NewServer(cfg.DBPath, cfg.ArtefactDir, cfg.AllowedImageFormats),
		siteConfig)

	fmt.Println("serving on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
