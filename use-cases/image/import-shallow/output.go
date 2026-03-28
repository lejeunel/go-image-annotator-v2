package import_shallow

type OutputPort interface {
	ErrNotFound(error)
	ErrInternal(error)
	ErrDependency(error)
	Success(Response)
}
