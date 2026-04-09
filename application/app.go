package app

import (
	af_store "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	im_store "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	"github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	sql_an "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/annotation"
	sql_clc "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	sql_im "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/image"
	sql_lbl "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
)

type SQLiteApp struct {
	ImageRepo      *sql_im.SQLiteImageRepo
	CollectionRepo *sql_clc.SQLiteCollectionRepo
	LabelRepo      *sql_lbl.SQLiteLabelRepo
	ImageStore     *im_store.ImageStore
	ArtefactRepo   *af_store.FileArtefactRepo
	AnnotationRepo *sql_an.SQLiteAnnotationRepo
}

func NewSQLiteApp(dbPath, artefactDir string) *SQLiteApp {
	db := sqlite.NewSQLiteDB(dbPath)
	imrepo := sql_im.NewSQLiteImageRepo(db)
	anrepo := sql_an.NewSQLiteAnnotationRepo(db)
	clrepo := sql_clc.NewSQLiteCollectionRepo(db)
	lbrepo := sql_lbl.NewSQLiteLabelRepo(db)
	afrepo := af_store.NewFileArtefactRepo(artefactDir)
	imstore := im_store.NewImageStore(imrepo, clrepo, anrepo, afrepo)
	return &SQLiteApp{
		ImageRepo:      imrepo,
		CollectionRepo: clrepo,
		LabelRepo:      lbrepo,
		ImageStore:     imstore,
		ArtefactRepo:   afrepo,
		AnnotationRepo: anrepo,
	}

}
