package remove

type OutputPort interface {
	Error(error)
	SuccessDeleteAnnotation(Response)
}
