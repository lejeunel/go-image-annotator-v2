package ingest

type OutputPort interface {
	Success(Response)
	ErrCollectionNotFound(error)
	ErrLabelNotFound(error)
	ErrInvalidImageData(error)
	ErrInternal(error)
	ErrDuplicateImage(error)
	ErrValidation(error)
}
