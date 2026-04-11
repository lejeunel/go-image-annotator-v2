package reader

import (
	"bytes"
	"errors"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"testing"

	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestDetectImageTypeFromNonImageBytesShouldFail(t *testing.T) {
	detector := ImageMIMETypeDetector{}
	_, _, err := detector.Detect(bytes.NewBuffer([]byte("asdf")))
	if !errors.Is(err, e.ErrValidation) {
		t.Fatalf("expected validation error, got %v", err)
	}
}

func TestDetectImageFromJPGBytes(t *testing.T) {
	detector := ImageMIMETypeDetector{}
	mimetype, _, _ := detector.Detect(bytes.NewBuffer(testJPGImage))
	if *mimetype != "image/jpeg" {
		t.Fatalf("expected mimetype image/jpg, got %v", mimetype)
	}
}

func TestDetectImageFromPNGBytes(t *testing.T) {
	detector := ImageMIMETypeDetector{}
	mimetype, _, _ := detector.Detect(bytes.NewBuffer(testPNGImage))
	if *mimetype != "image/png" {
		t.Fatalf("expected mimetype image/png, got %v", mimetype)
	}
}

func TestRecoverBytesAfterDetection(t *testing.T) {
	detector := ImageMIMETypeDetector{}
	_, reader, _ := detector.Detect(bytes.NewBuffer(testPNGImage))
	r, _ := io.ReadAll(reader)
	if !bytes.Equal(r, testPNGImage) {
		t.Fatal("expected to retrieve original bytes after detection")
	}

}
