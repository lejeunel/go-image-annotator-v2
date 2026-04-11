package interactors

import (
	clc "github.com/lejeunel/go-image-annotator-v2/use-cases/collection"
	im "github.com/lejeunel/go-image-annotator-v2/use-cases/image"
	lbl "github.com/lejeunel/go-image-annotator-v2/use-cases/label"
)

type Interactors struct {
	Label      *lbl.Interactors
	Collection *clc.Interactors
	Image      *im.Interactors
}
