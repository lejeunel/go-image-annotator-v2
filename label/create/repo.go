package create

type Repo interface {
	Create(Model) error
	Exists(string) (bool, error)
}
