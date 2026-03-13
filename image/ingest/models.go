package ingest

import (
	"io"
)

type Request struct {
	Collection string
	Labels     []string
	Reader     io.Reader
}

type Response struct {
	Collection string
}
