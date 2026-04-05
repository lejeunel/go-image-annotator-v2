package logging

import (
	"io"
	"log/slog"
)

func NewNoOpLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
