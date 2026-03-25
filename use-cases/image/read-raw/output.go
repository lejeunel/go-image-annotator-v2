package read_raw

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
}
