package add_bbox

import (
	"errors"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	imageStore im.ImageStore
	repo       Repo
}

func NewInteractor(imageStore im.ImageStore, repo Repo) *Interactor {
	return &Interactor{repo: repo, imageStore: imageStore}
}
func (i *Interactor) Execute(r Request, out OutputPort) {
	image, ok := i.findImage(r.ImageId, r.Collection, out)
	if !ok {
		return
	}

	label, ok := i.findLabel(r.Label, out)
	if !ok {
		return
	}

	box := a.NewBoundingBox(a.NewAnnotationId(), r.Xc, r.Yc, r.Width, r.Height, *label)
	if ok := i.validateBox(image, *box, out); !ok {
		return
	}

	if ok := i.addBox(image, *box, out); !ok {
		return
	}

	out.Success(Response{})

}
func (i *Interactor) addBox(image *im.Image, box a.BoundingBox, out OutputPort) bool {
	err := i.repo.AddBoundingBox(image.Id, image.Collection.Id, box)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(err)
			return false
		default:
			out.ErrInternal(e.ErrInternal)
			return false
		}
	}
	return true
}

func (i *Interactor) validateBox(image *im.Image, box a.BoundingBox, out OutputPort) bool {
	err := image.AddBoundingBox(box)
	if err != nil {
		out.ErrValidation(err)
		return false
	}
	return true
}

func (i *Interactor) findLabel(name string, out OutputPort) (*lbl.Label, bool) {
	label, err := i.repo.FindLabelByName(name)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(e.ErrNotFound)
			return nil, false
		default:
			out.ErrInternal(e.ErrInternal)
			return nil, false
		}
	}
	return label, true
}

func (i *Interactor) findImage(imageId im.ImageId, collectionName string, out OutputPort) (*im.Image, bool) {
	image, err := i.imageStore.Find(im.BaseImage{ImageId: imageId, Collection: collectionName})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(e.ErrNotFound)
			return nil, false
		default:
			out.ErrInternal(e.ErrInternal)
			return nil, false
		}
	}
	return image, true
}
