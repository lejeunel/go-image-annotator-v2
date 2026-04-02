package decoder

import (
	"bytes"
	"encoding/base64"
	"fmt"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
	i "image"
	_ "image/jpeg"
	_ "image/png"
	"slices"
)

type Base64ImageDecoder struct {
	AllowedFormats []string
}

func NewBase64ImageDecoder(allowedFormats []string) *Base64ImageDecoder {
	return &Base64ImageDecoder{AllowedFormats: allowedFormats}

}

// Decode decodes an image that has been encoded in base64.
// If the raw-bytes cannot be decoded to specified allowed formats,
// return an internal error
// The string returned is the format name used during format registration.
func (r *Base64ImageDecoder) Decode(data any) ([]byte, *string, error) {
	errCtx := "decoding base64 data"

	strData, ok := data.(string)
	if !ok {
		return nil, nil, fmt.Errorf("%v: expected string input, got %T: %w", errCtx, data, e.ErrImageFormat)
	}

	bytesData, err := base64.StdEncoding.DecodeString(strData)
	if err != nil {
		return nil, nil, fmt.Errorf("%v: %v: %w", errCtx, err, e.ErrImageFormat)
	}

	_, format, err := i.Decode(bytes.NewBuffer(bytesData))
	if err != nil {
		return nil, nil, fmt.Errorf("%v: decoding image data: %v: %w", errCtx, err, e.ErrImageFormat)
	}

	if !slices.Contains(r.AllowedFormats, format) {
		return nil, nil, fmt.Errorf("%v: checking for supported format (allowed formats are %v): %v: %w",
			errCtx, r.AllowedFormats, err, e.ErrImageFormat)
	}

	return bytesData, &format, nil
}
