package modify_bbox

type OutputPort interface {
	Error(error)
	Success(Response)
}
