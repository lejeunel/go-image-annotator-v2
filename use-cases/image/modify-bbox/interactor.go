package modify_bbox

import (
	"errors"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	output OutputPort
	repo   Repo
}

func NewInteractor(output OutputPort, repo Repo) *Interactor {
	return &Interactor{output: output, repo: repo}
}
func (i *Interactor) Execute(r Request) {
	label, ok := i.findLabel(r.Label)
	if !ok {
		return
	}
	u, ok := i.validate(r.AnnotationId, r.Xc, r.Yc, r.Width, r.Height, *label)
	if !ok {
		return
	}

	if ok := i.update(*u); !ok {
		return
	}
	i.output.Success(Response{})

}

func (i *Interactor) update(u Updatables) bool {
	if err := i.repo.UpdateBoundingBox(u); err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			i.output.ErrNotFound(e.ErrNotFound)
			return false
		default:
			i.output.ErrInternal(e.ErrInternal)
			return false
		}
	}
	return true

}

func (i *Interactor) validate(id a.AnnotationId, xc float32, yc float32, width float32,
	height float32, label lbl.Label) (*Updatables, bool) {

	if err := a.ValidateBoundingBox(xc, yc, width, height); err != nil {
		i.output.ErrValidation(err)
		return nil, false
	}
	return &Updatables{LabelId: label.Id, AnnotationId: id, Xc: xc, Yc: yc, Width: width, Height: height}, true

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
