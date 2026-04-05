package remove_bbox

type OutputPort interface {
	Error(error)
	Success(Response)
}
