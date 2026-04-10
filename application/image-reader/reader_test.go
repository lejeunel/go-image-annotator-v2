package reader

import (
	"bytes"
	_ "embed"
	"errors"
	"io"
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestImageReaderShouldFailOnArtefactRepoErrror(t *testing.T) {
	reader := NewImageReader(im.NewImageId(), &FakeReadArtefactRepo{Err: e.ErrInternal})
	_, err := io.ReadAll(reader)
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestImageReaderReturnsCorrectBytes(t *testing.T) {
	reader := NewImageReader(im.NewImageId(), &FakeReadArtefactRepo{Data: testJPGImage})
	r, _ := io.ReadAll(reader)
	if !bytes.Equal(r, testJPGImage) {
		t.Fatalf("did not retrieve original bytes")
	}
}
