package scroller

import (
	"errors"
	"fmt"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	e "github.com/lejeunel/go-image-annotator-v2/shared/errors"
)

type Scroller struct {
	currentImage im.ImageId
	repo         Repo
	criteria     ScrollingCriteria
}

type ScrollerState struct {
	Next     *im.BaseImage
	Previous *im.BaseImage
}

func (s *Scroller) getOne(direction ScrollingDirection) (*im.BaseImage, error) {

	image, err := s.repo.GetAdjacent(s.currentImage, s.criteria, direction)
	if err != nil && !errors.Is(err, e.ErrNotFound) {
		return nil, err
	}
	return image, nil
}

func (s *Scroller) State() (*ScrollerState, error) {
	state := ScrollerState{}
	next, errNext := s.getOne(ScrollNext)
	prev, errPrev := s.getOne(ScrollPrevious)

	if errNext != nil || errPrev != nil {
		return nil, fmt.Errorf("%w, %w", errNext, errPrev)
	}
	state.Next = next
	state.Previous = prev

	return &state, nil

}

func checkCriteria(repo Repo, imageId im.ImageId, criteria ScrollingCriteria) error {
	errCtx := "initializing image scroller"
	if err := repo.ImageMustExist(imageId); err != nil {
		return fmt.Errorf("%v: checking that image with id %v exists: %w",
			errCtx, imageId, err)

	}
	if criteria.Collection != nil {
		if err := repo.CollectionMustExist(*criteria.Collection); err != nil {
			return fmt.Errorf("%v: checking that collection with name %v exists: %w",
				errCtx, *criteria.Collection, err)
		}
	}
	return nil
}

func New(repo Repo, imageId im.ImageId, opts ...Option) (*Scroller, error) {
	criteria := NewCriteria(opts...)
	if err := checkCriteria(repo, imageId, criteria); err != nil {
		return nil, err
	}
	return &Scroller{repo: repo, currentImage: imageId, criteria: criteria}, nil
}
