package read_raw

type OutputPort interface {
	Success(Response)
	Error(error)
}
