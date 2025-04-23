// v0.2.0
// Author: wunderbarb
// (c), Apr 2025

package imgpath

import (
	_ "image/png"
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
	require.Equal(90, a)
	_, _, ok = c.ContinuousDarker(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 10,
	})
	// require.False(ok)

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
	require.Equal(90, a)
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

func TestImagePath_ContinuousDarkerExact(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test1.png")
	s, ok := c.ContinuousDarkerExact(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 9,
	})
	require.True(ok)
	require.Equal(uint8(0x5f), s)
	_, ok = c.ContinuousDarkerExact(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 6,
	})
	require.False(ok)
}

func TestImagePath_ContinuousBrighterExact(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test3.png")

	s, ok := c.ContinuousBrighterExact(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 9,
	})
	require.True(ok)
	require.Equal(uint8(0xC5), s)
	_, ok = c.ContinuousBrighterExact(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 5,
	})
	require.False(ok)
}

func TestImagePath_Continuous(t *testing.T) {
	require, _ := test.Describe(t)

	c := initC3("test1.png")
	co, ok := c.Continuous(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 9,
	})
	require.True(ok)
	require.True(co.Darker)
	c = initC3("test3.png")
	s, ok := c.Continuous(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 9,
	})
	require.True(ok)
	require.False(s.Darker)
	_, ok = c.Continuous(ContinuousInput{
		X:      2,
		Y:      2,
		T:      0x20,
		Length: 10,
	})
	require.False(ok)
}

func TestImagePath_AllBrighter(t *testing.T) {
	require, _ := test.Describe(t)

	c2 := initC2("test1.png")
	c2.SetCenter(2, 2)
	require.True(c2.AllBrighter(0x80))
	c3 := initC3("test1.png")
	require.False(c3.AllBrighter(0x80))
}
