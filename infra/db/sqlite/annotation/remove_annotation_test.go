package sqlite

import (
	"errors"
	a "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
	"testing"
)

func TestInternalErrOnRemoveImageLabelShouldFail(t *testing.T) {
	repos := NewAnnotationTestRepos()
	image, collection, label := CreateAnnotableImage(repos, "a-collection", "a-label")
	annotationId := a.NewAnnotationId()
	repos.Annotation.AddImageLabel(annotationId, image.Id, collection.Id, label.Id)
	repos.Annotation.Db.Close()
	err := repos.Annotation.RemoveAnnotation(annotationId)
	if !errors.Is(err, e.ErrInternal) {
		t.Fatalf("expected internal error, got %v", err)
	}
}

func TestRemoveAnnotation(t *testing.T) {
	repos := NewAnnotationTestRepos()
	image, collection, label := CreateAnnotableImage(repos, "a-collection", "a-label")
	annotationId := a.NewAnnotationId()
	repos.Annotation.AddImageLabel(annotationId, image.Id, collection.Id, label.Id)
	err := repos.Annotation.RemoveAnnotation(annotationId)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
