package annotator

import (
	sto "github.com/lejeunel/go-image-annotator-v2/application/image-store"
	"github.com/lejeunel/go-image-annotator-v2/application/scroller"
	an "github.com/lejeunel/go-image-annotator-v2/entities/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	addbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/add-bbox"
	updbox "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/modify-bbox"
	del "github.com/lejeunel/go-image-annotator-v2/use-cases/annotate/remove"
)

type Annotator struct {
	scroller   scroller.Interface
	store      sto.Interface
	boxAdder   addbox.Interface
	boxUpdater updbox.Interface
	boxDeleter del.Interface
}

func (a *Annotator) DeleteBox(r del.Request, view AnnotatorView) {
	a.boxDeleter.Execute(r, view)
}

func (a *Annotator) UpdateBox(r updbox.Request, view AnnotatorView) {
	a.boxUpdater.Execute(r, view)
}

func (a *Annotator) AddBox(r addbox.Request, view AnnotatorView) {
	a.boxAdder.Execute(r, view)
}

func (a *Annotator) Init(imageId im.ImageId, collection string, view AnnotatorView) {
	scrollerState, err := a.scroller.Init(imageId, scroller.WithCollection(collection))
	if err != nil {
		view.Error(err)
		return
	}
	view.DrawScroller(*scrollerState)
	view.DrawImageInfo(ImageInfo{Id: imageId, Collection: collection})
	a.drawImage(imageId, collection, view)
	a.drawAnnotationList(imageId, collection, view)

}
func (a *Annotator) drawImage(imageId im.ImageId, collection string, view AnnotatorView) {
	image, err := a.store.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	if err != nil {
		view.Error(err)
		return
	}
	view.DrawImage(*image)
}

func (a *Annotator) drawAnnotationList(imageId im.ImageId, collection string, view AnnotatorView) {
	image, err := a.store.Find(im.BaseImage{ImageId: imageId, Collection: collection})
	if err != nil {
		view.Error(err)
		return
	}
	annotations := []an.Annotation{}
	for _, box := range image.BoundingBoxes {
		annotations = append(annotations, an.Annotation{Id: box.Id, Label: box.Label.Name})
	}
	view.DrawAnnotationList(annotations)
}

func NewAnnotator(scrollerRepo scroller.Repo, store *sto.ImageStore) *Annotator {
	return &Annotator{scroller: scroller.New(scrollerRepo),
		store: store}
}
