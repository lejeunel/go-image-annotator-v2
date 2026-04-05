package list

type OutputPort interface {
	Success(Response)
	ErrInternal(error)
}
