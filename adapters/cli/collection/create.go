package collection

import (
	"fmt"
	"github.com/lejeunel/go-image-annotator-v2/config"

	cli "github.com/lejeunel/go-image-annotator-v2/adapters/cli"
	app "github.com/lejeunel/go-image-annotator-v2/application"
	clc "github.com/lejeunel/go-image-annotator-v2/use-cases/collection/create"
)

type CreatePresenter struct {
	cli.ErrorPresenter
}

func (p CreatePresenter) Success(r clc.Response) {
	fmt.Println("created collection with name", r.Name, "and description", r.Description)
}
func (p CreatePresenter) Error(err error) {
	fmt.Println(err.Error())
}

func Create(name, description string) {
	cfg := config.Parse()
	app := app.NewSQLiteApp(cfg.DBPath, cfg.ArtefactDir)
	itr := clc.NewDefaultInteractor(app.CollectionRepo)
	itr.Execute(clc.Request{Name: name, Description: description}, CreatePresenter{})

}
