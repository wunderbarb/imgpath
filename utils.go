// v0.1.0

package imgpath

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// CountContinuousBitsAtOne returns the maximum number of consecutive 1s in a `l`-bit word and the index of the first
// 1.
func CountContinuousBitsAtOne(n uint32, l int) (start int, length int) {
	nn := uint64(n) + (uint64(n) << l)
	var startI, len int
	var continuous bool
	i := 0
	for {
		switch nn&1 == 1 {
		case true:
			len++
			if !continuous {
				startI = i
				continuous = true
			}
		case false:
			if continuous {
				if len > length {
					length = len
					start = startI
				}
				len = 0
				continuous = false
			}
		}
		i++
		nn >>= 1
		if !continuous && i >= l {
			return
		}
	}
}

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
