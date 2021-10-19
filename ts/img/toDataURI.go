package img

import (
	"fmt"
	"image"
)

func (p *Img) MustToDataURI() (data string) {
	data, err := ToDataURI(p.Image, p.Type)
	if err != nil {
		panic(err)
	}
	return
}

func (p *Img) ToDataURI() (data string, err error) {
	return ToDataURI(p.Image, p.Type)
}

// ToDataURI  支持 data uri
func ToDataURI(target image.Image, imageType string) (data string, err error) {
	data, err = ToBase64(target, imageType)
	if err != nil {
		return "", err
	}
	data = fmt.Sprintf("data:image/%s;base64,%s", imageType, data)
	return
}
