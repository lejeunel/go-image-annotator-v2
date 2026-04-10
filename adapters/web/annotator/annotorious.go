package annotator

import (
	"bytes"
	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"text/template"
)

type AnnotatorState struct {
	ImageId          string
	Collection       string
	Annotations      string
	EnableAnnotation bool
	OriginType       string
	OriginId         string
	Ordering         string
	Descending       bool
}

func MakeAnnotoriousScript(imageId im.ImageId, collection string) (*Node, error) {
	tAnnot, err := template.New("annotator").ParseFS(templatesFiles, "templates/annotator.js")
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBufferString("")
	data := AnnotatorState{ImageId: imageId.String(),
		Collection: collection, Annotations: "[]",
		EnableAnnotation: true}

	err = tAnnot.ExecuteTemplate(buf, "annotator", data)
	if err != nil {
		return nil, err
	}
	script := Script(Raw(buf.String()))
	return &script, nil
}
