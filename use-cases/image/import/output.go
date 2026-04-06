package import_image

type OutputPort interface {
	Error(error)
	Success(Response)
}
