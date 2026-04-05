package cmd

import (
	"fmt"
	api "github.com/lejeunel/go-image-annotator-v2/adapters/api"
	apiServer "github.com/lejeunel/go-image-annotator-v2/adapters/api/server"
	"github.com/lejeunel/go-image-annotator-v2/config"
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
	apiServer := apiServer.NewServer(cfg.DBPath, cfg.ArtefactDir, cfg.AllowedImageFormats)
	api.RegisterAPI(mux, *apiServer)
	fmt.Println("serving on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
