package delete

type Repo interface {
	Exists(string) (bool, error)
	Delete(string) error
	IsUsed(string) (bool, error)
}
