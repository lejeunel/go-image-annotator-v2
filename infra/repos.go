package infra

import (
	af_store "github.com/lejeunel/go-image-annotator-v2/application/artefact-store"
	im_store "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	"github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite"
	an "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/collection"
	im "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/infra/db/sqlite/label"
)

type SQLiteInfra struct {
	ImageRepo      *im.SQLiteImageRepo
	CollectionRepo *clc.SQLiteCollectionRepo
	LabelRepo      *lbl.SQLiteLabelRepo
	ImageStore     *im_store.ImageStore
	ArtefactRepo   *af_store.FileArtefactRepo
	AnnotationRepo *an.SQLiteAnnotationRepo
}

func NewSQLiteInfra(dbPath, artefactDir string) *SQLiteInfra {
	db := sqlite.NewSQLiteDB(dbPath)
	imrepo := im.NewSQLiteImageRepo(db)
	anrepo := an.NewSQLiteAnnotationRepo(db)
	clrepo := clc.NewSQLiteCollectionRepo(db)
	lbrepo := lbl.NewSQLiteLabelRepo(db)
	afrepo := af_store.NewFileArtefactRepo(artefactDir)
	imstore := im_store.NewImageStore(imrepo, clrepo, anrepo, afrepo)
	return &SQLiteInfra{
		ImageRepo:      imrepo,
		CollectionRepo: clrepo,
		LabelRepo:      lbrepo,
		ImageStore:     imstore,
		ArtefactRepo:   afrepo,
		AnnotationRepo: anrepo,
	}

}
