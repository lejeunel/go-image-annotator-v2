package ingest

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrInternal(error)
	ErrDuplication(error)
	ErrValidation(error)
}
