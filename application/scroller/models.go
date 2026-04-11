package scroller

import (
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type ScrollingDirection int

const (
	ScrollNext ScrollingDirection = iota
	ScrollPrevious
)

type Request struct {
	ImageId    im.ImageId
	Collection string
	Direction  ScrollingDirection
}
