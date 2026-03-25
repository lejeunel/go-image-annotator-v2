package image

type ArtefactRepo interface {
	Store(ImageId, []byte) error
	Delete(ImageId) error
	Get(ImageId) ([]byte, error)
}

type FakeArtefactRepo struct {
	GotArtefact      bool
	Err              error
	NumDeletedImages int
	Data             []byte
}

func (r *FakeArtefactRepo) Store(ImageId, []byte) error {
	if r.Err != nil {
		return r.Err
	}
	r.GotArtefact = true
	return nil
}

func (r *FakeArtefactRepo) Delete(ImageId) error {
	r.NumDeletedImages += 1
	return nil
}

func (r *FakeArtefactRepo) Get(ImageId) ([]byte, error) {
	return r.Data, nil
}
