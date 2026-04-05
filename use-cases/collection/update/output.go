package update

type OutputPort interface {
	Success(Response)
	ErrDuplication(error)
	ErrNotFound(error)
	ErrInternal(error)
}
