package delete

type OutputPort interface {
	ErrDependency(error)
	ErrInternal(error)
	ErrNotFound(error)
	Success()
}
