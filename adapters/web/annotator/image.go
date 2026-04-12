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

func (p *ImageView) Render(image im.Image) Node {
	if image.Reader == nil {
		return Text("presenting image: got no reader")

	}
	bytes, err := io.ReadAll(image.Reader)

	if err != nil {
		return Text(err.Error())
	}

	b64Image := base64.StdEncoding.EncodeToString(bytes)
	return Img(ID("image"), Src(fmt.Sprintf("data:%v;base64,%s",
		image.MIMEType, b64Image)))

}
