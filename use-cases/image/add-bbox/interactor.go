package add_bbox

import (
	"errors"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	im "github.com/lejeunel/go-image-annotator-v2/domain/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output     OutputPort
	imageStore im.ImageStore
	repo       Repo
}

func NewInteractor(output OutputPort, imageStore im.ImageStore, repo Repo) *Interactor {
	return &Interactor{output: output, repo: repo, imageStore: imageStore}
}
func (i *Interactor) Execute(r Request) {
	image, ok := i.findImage(r.ImageId, r.Collection)
	if !ok {
		return
	}

	label, ok := i.findLabel(r.Label)
	if !ok {
		return
	}

	box := a.NewBoundingBox(a.NewAnnotationId(), r.Xc, r.Yc, r.Width, r.Height, *label)
	if ok := i.validateBox(image, *box); !ok {
		return
	}

	if ok := i.addBox(image, *box); !ok {
		return
	}

	i.output.Success(Response{})

}
func (i *Interactor) addBox(image *im.Image, box a.BoundingBox) bool {
	err := i.repo.AddBoundingBox(image.Id, image.Collection.Id, box)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(err)
			return false
		default:
			i.output.ErrInternal(e.ErrInternal)
			return false
		}
	}
	return true
}

func (i *Interactor) validateBox(image *im.Image, box a.BoundingBox) bool {
	err := image.AddBoundingBox(box)
	if err != nil {
		i.output.ErrValidation(err)
		return false
	}
	return true
}

func (i *Interactor) findLabel(name string) (*lbl.Label, bool) {
	label, err := i.repo.FindLabelByName(name)
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(e.ErrNotFound)
			return nil, false
		default:
			i.output.ErrInternal(e.ErrInternal)
			return nil, false
		}
	}
	return label, true
}

func (i *Interactor) findImage(imageId im.ImageId, collectionName string) (*im.Image, bool) {
	image, err := i.imageStore.Find(im.BaseImage{ImageId: imageId, Collection: collectionName})
	if err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(e.ErrNotFound)
			return nil, false
		default:
			i.output.ErrInternal(e.ErrInternal)
			return nil, false
		}
	}
	return image, true
}
