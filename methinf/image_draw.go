package methinf

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

//Image represent the image to be drawn
type Image struct {
	Width, Height int
	colour        uint8
}

//Bounds is used to get rectangular bound
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

//ColorModel is used to get the color of the image
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

//At is used to determine where to draw
func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.colour + uint8(x), img.colour + uint8(y), 255, 255}
}

// func main() {
// 	m := Image{100, 100, 128}
// 	pic.ShowImage(&m)
// }
