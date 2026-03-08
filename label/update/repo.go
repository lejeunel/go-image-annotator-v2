package update

type Repo interface {
	Update(Model) error
}
