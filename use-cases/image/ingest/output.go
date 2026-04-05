package ingest

type OutputPort interface {
	Success(Response)
	Error(error)
}
