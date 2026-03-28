package modify_bbox

import (
	"errors"

	a "github.com/lejeunel/go-image-annotator-v2/domain/annotation"
	lbl "github.com/lejeunel/go-image-annotator-v2/domain/label"
	e "github.com/lejeunel/go-image-annotator-v2/errors"
)

type Interactor struct {
	repo Repo
}

func NewInteractor(repo Repo) *Interactor {
	return &Interactor{repo: repo}
}
func (i *Interactor) Execute(r Request, out OutputPort) {
	label, ok := i.findLabel(r.Label, out)
	if !ok {
		return
	}
	u, ok := i.validate(r.AnnotationId, r.Xc, r.Yc, r.Width, r.Height, *label, out)
	if !ok {
		return
	}

	if ok := i.update(*u, out); !ok {
		return
	}
	out.Success(Response{})

}

func (i *Interactor) update(u Updatables, out OutputPort) bool {
	if err := i.repo.UpdateBoundingBox(u); err != nil {
		switch {
		case errors.Is(err, e.ErrNotFound):
			out.ErrNotFound(e.ErrNotFound)
			return false
		default:
			out.ErrInternal(e.ErrInternal)
			return false
		}
	}
	return true

}

func (i *Interactor) validate(id a.AnnotationId, xc float32, yc float32, width float32,
	height float32, label lbl.Label, out OutputPort) (*Updatables, bool) {

	if err := a.ValidateBoundingBox(xc, yc, width, height); err != nil {
		out.ErrValidation(err)
		return nil, false
	}
	return &Updatables{LabelId: label.Id, AnnotationId: id, Xc: xc, Yc: yc, Width: width, Height: height}, true

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
