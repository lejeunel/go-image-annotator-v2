package update

type OutputPort interface {
	Success(Response)
	ErrDuplication(string)
	ErrNotFound(string)
	ErrInternal(string)
}
