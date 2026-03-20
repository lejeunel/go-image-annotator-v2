package assign_label

type OutputPort interface {
	Success(Response)
	ErrNotFound(error)
	ErrImageNotInCollection(error)
	ErrInternal(error)
}
