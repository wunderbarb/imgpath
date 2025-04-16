// V0.1.0

package imgpath

import (
	"path/filepath"
	"testing"

	"github.com/wunderbarb/test"
)

func TestNew(t *testing.T) {
	require, _ := test.Describe(t)

	var path []Pos
	_, err := New(path)
	require.ErrorIs(err, ErrNoPath)
	path = []Pos{{1, 1}}
	ip, err := New(path)
	require.NoError(err)
	require.NotNil(ip)
}

func TestImagePath_Next(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test1.png")
	tests := []uint8{32, 128, 32, 32, 32, 32, 32, 32, 255, 255, 255, 32, 32}
	for _, tt := range tests {
		require.Equal(tt, c.Next())
	}
}

func TestImagePath_All(t *testing.T) {
	require, _ := test.Describe(t)

	var cnt = 0
	c := initC3("test1.png")

	c.All(func(_ uint8, _ int) {
		cnt++
	})
	require.Equal(c.Len(), cnt)
}

func TestImagePath_Until(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test1.png")

	cnt := 0
	c.Until(func(v uint8, index int) bool {
		if v != 255 {
			return true
		}
		cnt = index
		return false
	})
	require.Equal(8, cnt)
}

func TestImagePath_Threshold(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test1.png")

	nn := c.Diff()
	require.Len(nn, C3.Len())
	require.Equal([]int{-223, -127, -223, -223, -223, -223, -223, -223, 0, 0, 0, -223}, nn)
}

// -----------------

func initC3(file string) ImagePath {
	img, err := GrayFromFile(filepath.Join("testfixtures", file))
	isPanic(err)
	c := C3
	c.SetImage(img)
	return c
}

func initC4(file string) ImagePath {
	img, err := GrayFromFile(filepath.Join("testfixtures", file))
	isPanic(err)
	c := C4
	c.SetImage(img)

	return c
}

func initC5(file string) ImagePath {
	img, err := GrayFromFile(filepath.Join("testfixtures", file))
	isPanic(err)
	c := C5
	c.SetImage(img)
	return c
}

func isPanic(err error) {
	if err != nil {
		panic(err)
	}
}
