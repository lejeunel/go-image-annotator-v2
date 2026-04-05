package add_bbox

type OutputPort interface {
	Error(error)
	Success(Response)
}
