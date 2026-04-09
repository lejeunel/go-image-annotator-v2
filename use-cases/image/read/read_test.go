package read

import (
	"testing"

	st "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	clc "github.com/lejeunel/go-image-annotator-v2/entities/collection"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/entities/label"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

func TestHandleNotFoundError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&st.FakeImageStore{Err: e.ErrNotFound})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if !p.GotNotFoundErr || p.GotSuccess {
		t.Fatalf("expected to get not found error")
	}
}

func TestHandleInternalError(t *testing.T) {
	p := &FakePresenter{}
	itr := NewInteractor(&st.FakeImageStore{Err: e.ErrInternal})
	itr.Execute(Request{ImageId: im.NewImageId(), Collection: "a-collection"}, p)
	if !p.GotInternalErr || p.GotSuccess {
		t.Fatalf("expected to get internal error")
	}
}

func TestFindImage(t *testing.T) {
	p := &FakePresenter{}
	existingImage := im.NewImage(im.NewImageId(), *clc.NewCollection(clc.NewCollectionId(), "a-collection"))
	label := lbl.NewLabel(lbl.NewLabelId(), "a-label")
	existingImage.AddLabel(label)
	existingImage.AddBoundingBox(*a.NewBoundingBox(a.NewAnnotationId(), 1, 1, 1, 1, *label))
	itr := NewInteractor(&st.FakeImageStore{Return: existingImage})
	itr.Execute(Request{ImageId: existingImage.Id, Collection: existingImage.Collection.Name}, p)
	got := p.Got
	if !p.GotSuccess {
		t.Fatalf("expected to get success")
	}
	if !(got.Id == existingImage.Id) || !(got.Collection.Name == existingImage.Collection.Name) {
		t.Fatalf("expected to get image id: %v, collection %v, got %v, %v",
			existingImage.Id, existingImage.Collection.Name,
			got.Id, got.Collection)

	}
	if (len(got.Labels) != 1) || (len(got.BoundingBoxes) != 1) {
		t.Fatalf("expected to get num. labels %v and num. boxes %v, got %v, %v",
			len(existingImage.Labels), len(existingImage.BoundingBoxes),
			len(got.Labels), len(got.BoundingBoxes))

	}
}
