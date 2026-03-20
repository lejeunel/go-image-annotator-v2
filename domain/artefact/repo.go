package artefact

type ArtefactRepo interface {
	Store(ArtefactId, []byte) error
	Delete(ArtefactId) error
}

type FakeArtefactRepo struct {
	GotArtefact      bool
	Err              error
	NumDeletedImages int
}

func (r *FakeArtefactRepo) Store(ArtefactId, []byte) error {
	if r.Err != nil {
		return r.Err
	}
	r.GotArtefact = true
	return nil
}

func (r *FakeArtefactRepo) Delete(ArtefactId) error {
	r.NumDeletedImages += 1
	return nil
}
