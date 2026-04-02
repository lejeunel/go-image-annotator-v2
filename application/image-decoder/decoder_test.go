package decoder

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"errors"
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

//go:embed sample-image.jpg
var testJPGImage []byte

func TestErrOnInvalidDataShouldFail(t *testing.T) {
	decoder := NewBase64ImageDecoder([]string{"jpeg"})
	_, _, err := decoder.Decode("invalid-data")
	if !errors.Is(err, e.ErrImageFormat) {
		t.Fatalf("expected image format error on invalid base64 data, got %v", err)
	}
}

func TestDecodeJPGImage(t *testing.T) {
	decoder := NewBase64ImageDecoder([]string{"jpeg"})
	_, format, err := decoder.Decode(base64.StdEncoding.EncodeToString(testJPGImage))
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}
	if *format != "jpeg" {
		t.Fatalf("expected jpeg format, got %v", *format)
	}
}

func TestFormatNotAllowedShouldFail(t *testing.T) {
	decoder := NewBase64ImageDecoder([]string{"png"})
	_, _, err := decoder.Decode(base64.StdEncoding.EncodeToString(testJPGImage))
	if !errors.Is(err, e.ErrImageFormat) {
		t.Fatalf("expected error on invalid format got %v", err)
	}
}

func TestNonStringInputShouldFail(t *testing.T) {
	decoder := NewBase64ImageDecoder([]string{"png"})
	_, _, err := decoder.Decode([]byte("asdf"))
	if !errors.Is(err, e.ErrImageFormat) {
		t.Fatalf("expected error on invalid input got %v", err)
	}
}

func RecoverEncodedBytes(t *testing.T) {
	decoder := NewBase64ImageDecoder([]string{"jpeg"})
	data, _, err := decoder.Decode(base64.StdEncoding.EncodeToString(testJPGImage))
	if err != nil {
		t.Fatalf("did not expect error, got %v", err)
	}
	if !bytes.Equal(data, testJPGImage) {
		t.Fatal("did not recover original byte slice")
	}

}
