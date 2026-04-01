package cmd

import (
	"fmt"
	api "github.com/lejeunel/go-image-annotator-v2/api"
	apiServer "github.com/lejeunel/go-image-annotator-v2/api/server"
	"github.com/lejeunel/go-image-annotator-v2/config"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	port     int
	migrate  bool
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
	ServeCmd.Flags().BoolVarP(&migrate, "migrate", "m", false, "apply DB migrations")
}

func serve(port int) {
	cfg := config.Parse()
	mux := http.NewServeMux()
	apiServer := apiServer.NewServer(cfg.DBPath)
	api.RegisterAPI(mux, *apiServer)
	fmt.Println("serving on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
