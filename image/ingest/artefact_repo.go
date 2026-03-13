package ingest

type ArtefactRepo interface {
	Store([]byte) error
}

type FakeErrArtefactRepo struct {
	err error
}

func (r *FakeErrArtefactRepo) Store(data []byte) error {
	return r.err
}

type FakeArtefactRepo struct {
}

func (r *FakeArtefactRepo) Store(data []byte) error {
	return nil
}
