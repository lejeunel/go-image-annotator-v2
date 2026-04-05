package create

type OutputPort interface {
	Success(Response)
	ErrDuplication(error)
	ErrInternal(error)
	ErrValidation(error)
}
