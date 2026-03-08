package delete

type Repo interface {
	Delete(string) error
	IsUsed(string) (bool, error)
}
