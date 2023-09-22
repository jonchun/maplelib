package wz

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

// A B64MapleCanvas is a wrapper around image.Image that holds info
// about a png image extracted from a wz file
type B64MapleCanvas struct {
	width  int
	height int
	img    *image.Image
}

// NewB64MapleCanvas initializes a new B64MapleCanvas object
// with the given file path and size
func NewB64MapleCanvas(w, h int, b64String string,
) *B64MapleCanvas {

	b, err := base64.StdEncoding.DecodeString(b64String)
	if err != nil {
		return nil
	}
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return nil
	}

	return &B64MapleCanvas{
		width:  w,
		height: h,
		img:    &img,
	}
}

func (f *B64MapleCanvas) Height() int { return f.height }
func (f *B64MapleCanvas) Width() int  { return f.width }
func (f *B64MapleCanvas) Image() *image.Image {
	return f.img
}
