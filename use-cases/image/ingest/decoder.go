package ingest

type ImageDecoder interface {
	Decode(data any) ([]byte, *string, error)
}
