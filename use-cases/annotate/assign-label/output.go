package assign_label

type OutputPort interface {
	Success(Response)
	Error(error)
}
