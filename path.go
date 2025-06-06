// v0.2.0
// Author: wunderbarb
// © April 2025

// Package imgpath manages operations on a path of pixels on a gray-scaled image.
package imgpath

import (
	"errors"
	"image"
)

// ErrNoPath occurs when creating an ImagePath with no defined path.
var ErrNoPath = errors.New("no path")

// Pos is a position in the path relative to the center.
type Pos struct {
	// X is the horizontal coordinate.
	X int
	// Y is the vertical coordinate.
	Y int
}

// ImagePath is the iterator of all pixels in the path.
type ImagePath struct {
	path        []Pos
	index       int
	img         *image.Gray
	centerPoint Pos
	cycled      bool
}

// New creates an ImagePath for the path `ps`.
func New(ps []Pos) (*ImagePath, error) {
	if len(ps) == 0 {
		return nil, ErrNoPath
	}
	var ip ImagePath
	ip.path = make([]Pos, len(ps))
	copy(ip.path, ps)
	return &ip, nil
}

// All method iterates function `fn` over the complete path.
func (ip ImagePath) All(fn func(v uint8, index int)) {
	for i := 0; i < len(ip.path); i++ {
		j := ip.index
		fn(ip.Next(), j)
	}
}

// Current returns the current position of the index relative to the center point.
func (ip ImagePath) Current() (int, int) {
	return ip.path[ip.index].X, ip.path[ip.index].Y
}

// CurrentAbsolute returns the absolute position of the current index in the image.
func (ip ImagePath) CurrentAbsolute() (int, int) {
	return ip.centerPoint.X + ip.path[ip.index].X, ip.centerPoint.Y + ip.path[ip.index].Y
}

// Cycled is true if a full path has been explored.
func (ip ImagePath) Cycled() bool {
	return ip.cycled
}

// Len returns the number of pixels in the path.
func (ip ImagePath) Len() int {
	return len(ip.path)
}

// Next returns the value of the next pixel.  It starts with position 0 and cycles.
func (ip *ImagePath) Next() uint8 {
	v := ip.img.GrayAt(ip.CurrentAbsolute()).Y
	ip.index++
	if ip.index >= len(ip.path) {
		ip.index = 0
		ip.cycled = true
	}
	return v
}

// Reset resets the path to the first pixel.
func (ip *ImagePath) Reset() {
	ip.index = 0
	ip.cycled = false
}

// SetCenter defines the reference point of the path.  It resets the cycle.
func (ip *ImagePath) SetCenter(x int, y int) {
	ip.centerPoint = Pos{x, y}
	ip.Reset()
}

// SetImage defines the image that will be iterated.
func (ip *ImagePath) SetImage(img image.Image) {
	ig, ok := img.(*image.Gray)
	if !ok {
		ig = Image2Gray(img)
	}
	ip.img = ig
}

// SetPath updates the path with the same center and image.
func (ip *ImagePath) SetPath(p []Pos) {
	ip.path = p
	ip.Reset()
}

// Until iterates over the path until the function `fn` returns false.
func (ip ImagePath) Until(fn func(v uint8, index int) bool) {
	for {
		j := ip.index
		if !fn(ip.Next(), j) {
			break
		}
	}
}

// Diff provides a slice of the difference between the path and the center pixel.  The first value corresponds to the
// first pixel of the path.
func (ip ImagePath) Diff() []int {
	nn := make([]int, len(ip.path))
	vv := ip.AtCenter()
	ip.All(func(v uint8, index int) {
		nn[index] = int(v) - int(vv)

	})
	return nn
}


// Than method provides a slice of the difference between the path and the threshold `t`.  The first value corresponds
// to the first pixel of the path.
func (ip ImagePath) Than(t uint8) []int {
	nn := make([]int, len(ip.path))
	ip.All(func(v uint8, index int) {
		nn[index] = int(v) - int(t)

	})
	return nn
}

// --------------------------

func (ip ImagePath) AtCenter() uint8 {
	return ip.img.GrayAt(ip.centerPoint.X, ip.centerPoint.Y).Y
}
