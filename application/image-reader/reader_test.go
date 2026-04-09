package reader

import (
	"bytes"
	_ "embed"
	"io"
	"testing"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

func TestImageReader(t *testing.T) {
	reader := NewImageReader(im.NewImageId(), &FakeReadArtefactRepo{Data: testJPGImage})
	r, _ := io.ReadAll(reader)
	if !bytes.Equal(r, testJPGImage) {
		t.Fatalf("did not retrieve original bytes")
	}
}
