package assign_label

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrDependency(error)
	ErrInternal(error)
}
