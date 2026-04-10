package reader

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func formatToMIME(format string) string {
	switch format {
	case "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "gif":
		return "image/gif"
	case "webp":
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}

type ImageMIMETypeDetector struct{}

func (d ImageMIMETypeDetector) Detect(r io.Reader) (*string, io.Reader, error) {
	// Read a small prefix (DecodeConfig does not need the full file)
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	_, format, err := image.DecodeConfig(tee)
	if err != nil {
		return nil, nil, fmt.Errorf("decoding image: %w: %w", err, e.ErrValidation)
	}

	// Reconstruct the full reader:
	// first the consumed bytes, then the remaining original reader
	newReader := io.MultiReader(&buf, r)

	mimetype := formatToMIME(format)
	return &mimetype, newReader, nil
}
