package modify_bbox

type OutputPort interface {
	ErrValidation(error)
	ErrNotFound(error)
	ErrInternal(error)
	Success(Response)
}
