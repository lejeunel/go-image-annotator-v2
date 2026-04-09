package image

import (
	"fmt"
	"github.com/lejeunel/go-image-annotator-v2/config"
	"os"
	"path/filepath"

	cli "github.com/lejeunel/go-image-annotator-v2/adapters/cli"
	ing "github.com/lejeunel/go-image-annotator-v2/use-cases/image/ingest"
)

type IngestPresenter struct {
	cli.ErrorPresenter
}

func (p *IngestPresenter) Success(r ing.Response) {
	fmt.Println("ingested image with id:", r.ImageId)
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
		itr.Execute(ing.Request{Collection: collection, Reader: f}, &IngestPresenter{})
	}

}
