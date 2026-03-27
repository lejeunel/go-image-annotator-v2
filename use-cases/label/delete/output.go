package delete

type OutputPort interface {
	ErrNotFound(error)
	ErrDependency(error)
	ErrInternal(error)
	Success()
}
