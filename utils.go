// v0.1.1
// Author: wunderbarb

package imgpath

import (
	"image"
	"image/draw"
	"os"
)

// Image2Gray converts `img` into a gray-scaled image.
func Image2Gray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	draw.Draw(gray, bounds, img, bounds.Min, draw.Src)
	return gray
}

// GrayFromFile reads an image from the `file` and converts it to a gray-scaled image.
func GrayFromFile(file string) (*image.Gray, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return Image2Gray(img), nil
}
