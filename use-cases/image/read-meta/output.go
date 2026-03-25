package read_meta

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
}
