package annotator

import (
	"encoding/base64"
	"fmt"
	"io"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type ImageView struct {
	result Node
}

func (p *ImageView) Success(image *im.Image) {
	if image.Reader == nil {
		p.result = Text("presenting image: got no reader")
		return

	}
	bytes, err := io.ReadAll(image.Reader)

	if err != nil {
		p.result = Text(err.Error())
		return
	}

	b64Image := base64.StdEncoding.EncodeToString(bytes)
	imNode := Img(ID("image"), Src(fmt.Sprintf("data:%v;base64,%s", image.MIMEType, b64Image)))
	p.result = imNode

}
func (p *ImageView) Error(err error) {
	p.result = Text(err.Error())
}

func (p *ImageView) Build() Node {
	return p.result
}
