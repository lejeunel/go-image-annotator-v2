package delete

type OutputPort interface {
	ErrNotFound(error)
	ErrInternal(error)
	Success(Response)
}
