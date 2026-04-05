package list

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
}
