package unassign_label

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
	ErrDependency(error)
}
