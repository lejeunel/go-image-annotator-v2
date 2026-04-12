package annotator

import (
	"github.com/lejeunel/go-image-annotator-v2/application/scroller"
	an "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	addbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/add-bbox"
	updbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/modify-bbox"
	del "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/remove"
)

type ImageInfo struct {
	Id         im.ImageId
	Collection string
}

type AddBoxView interface {
	Success(addbox.Response)
	Error(error)
}

type AnnotatorView interface {
	DrawScroller(scroller.ScrollerState)
	DrawImage(im.Image)
	DrawImageInfo(ImageInfo)
	DrawAnnotationList([]an.Annotation)
	Error(error)
	SuccessAddBox(addbox.Response)
	SuccessUpdateBox(updbox.Response)
	SuccessDeleteAnnotation(del.Response)
}
