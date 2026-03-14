package artefact

type ArtefactRepo interface {
	Store(ArtefactID, []byte) error
}

type FakeArtefactRepo struct {
	GotArtefact bool
	Err         error
}

func (r *FakeArtefactRepo) Store(ArtefactID, []byte) error {
	if r.Err != nil {
		return r.Err
	}
	r.GotArtefact = true
	return nil
}
