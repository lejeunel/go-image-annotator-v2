package remove_bbox

type OutputPort interface {
	ErrNotFound(error)
	ErrInternal(error)
	Success(Response)
}
