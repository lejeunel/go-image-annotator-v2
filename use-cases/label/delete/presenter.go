package delete

type OutputPort interface {
	ErrDependency(error)
	ErrInternal(error)
	Success()
}
