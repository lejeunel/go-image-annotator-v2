package import_shallow

type OutputPort interface {
	Error(error)
	Success(Response)
}
