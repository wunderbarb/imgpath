// v0.1.0

package imgpath

import (
	"fmt"
	"math"
	"path/filepath"
	"testing"

	"github.com/wunderbarb/test"
)

func TestImagePath_ContinuousDarker(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test1.png")

	s, a, ok := c.ContinuousDarker(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 6,
	})
	require.True(ok)
	require.Equal(uint8(0x5f), s)
	require.Equal(math.Pi/2, a)
	_, _, ok = c.ContinuousDarker(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 10,
	})
	require.False(ok)

	c1 := initC3("test2.png")
	_, _, ok = c1.ContinuousDarker(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x40,
		Length: 6,
	})
	require.False(ok)
}

func TestImagePath_ContinuousBrighter(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test3.png")

	s, a, ok := c.ContinuousBrighter(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 6,
	})
	require.True(ok)
	require.Equal(uint8(0xC5), s)
	require.Equal(math.Pi/2, a)
	_, _, ok = c.ContinuousBrighter(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 10,
	})
	require.False(ok)

	c1 := initC3("test4.png")
	_, _, ok = c1.ContinuousBrighter(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x40,
		Length: 6,
	})
	require.False(ok)
}

func TestImagePath_ContinuousDarkerAtLeast(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test1.png")

	s, ok := c.ContinuousDarkerAtLeast(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 6,
	})
	require.True(ok)
	require.Equal(uint8(0x5f), s)
	_, ok = c.ContinuousDarkerAtLeast(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 10,
	})
	require.False(ok)

	c1 := initC3("test2.png")
	_, ok = c1.ContinuousDarkerAtLeast(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x40,
		Length: 6,
	})
	require.False(ok)
}

func TestImagePath_ContinuousBrighterAtLeast(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test3.png")

	s, ok := c.ContinuousBrighterAtLeast(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 6,
	})
	require.True(ok)
	require.Equal(uint8(0xdd), s)
	_, ok = c.ContinuousBrighterAtLeast(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 10,
	})
	require.False(ok)

	c1 := initC3("test4.png")
	_, ok = c1.ContinuousBrighterAtLeast(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x40,
		Length: 6,
	})
	require.False(ok)
}

func TestCascadedFast(t *testing.T) {
	require, _ := test.Describe(t)

	img, err := GrayFromFile(filepath.Join("testfixtures", "test5.png"))
	isPanic(err)
	s, a, ok := CascadedFast(img, 4, 4, 32)
	require.True(ok)
	require.Equal(2*math.Pi*6/20.0, s)
	fmt.Println(a)
}
