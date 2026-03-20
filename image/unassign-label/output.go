package unassign_label

type OutputPort interface {
	ErrNotFound(error)
	ErrInternal(error)
}
