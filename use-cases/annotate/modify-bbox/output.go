package modify_bbox

type OutputPort interface {
	Error(error)
	SuccessUpdateBox(Response)
}
