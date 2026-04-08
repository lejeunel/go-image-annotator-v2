package cli

import (
	"fmt"
	"github.com/lejeunel/go-image-annotator-v2/config"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"

	ing "github.com/lejeunel/go-image-annotator-v2/use-cases/image/ingest"
)

var (
	IngestDirectoryCmd = &cobra.Command{
		Use:   "ingest-dir [dir] [collection]",
		Short: "Ingests all image located at [dir] directory into [collection]",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			dir := args[0]
			collection := args[1]
			fmt.Println("ingesting directory", dir, "into collection", collection)
			IngestDirectory(dir, collection)
		},
	}
)

type CLIIngestPresenter struct{}

func (p *CLIIngestPresenter) Success(r ing.Response) {
	fmt.Println("ingested image with id:", r.ImageId)
}
func (p *CLIIngestPresenter) Error(err error) {
	fmt.Println(err.Error())
}

func IngestDirectory(dir, collection string) {

	cfg := config.Parse()
	itr := ing.NewSQLiteIngestInteractor(cfg.DBPath, cfg.ArtefactDir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		ingestImage(itr, dir, entry, collection)
	}
}

func ingestImage(itr *ing.Interactor, dir string, entry os.DirEntry, collection string) {
	if !entry.IsDir() {
		f, err := os.Open(filepath.Join(dir, entry.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}
		itr.Execute(ing.Request{Collection: collection, Reader: f}, &CLIIngestPresenter{})
	}

}
