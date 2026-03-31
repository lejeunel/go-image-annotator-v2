package add_bbox

type OutputPort interface {
	ErrNotFound(error)
	ErrInternal(error)
	ErrValidation(error)
	Success(Response)
}
