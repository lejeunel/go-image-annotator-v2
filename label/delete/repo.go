package delete

type Repo interface {
	Delete(r Model) error
}
