// v0.1.0

package imgpath

import (
	"testing"

	"github.com/wunderbarb/test"
)

func TestIndexRing(t *testing.T) {
	require, _ := test.Describe(t)
	ri, err := NewIndexRing(3)
	require.NoError(err)
	require.Equal(0, ri.Next())
	require.False(ri.Cycled())
	require.Equal(1, ri.Next())
	require.False(ri.Cycled())
	require.Equal(2, ri.Next())
	require.True(ri.Cycled())
	require.Equal(0, ri.Next())

}

func Test_continuousDark(t *testing.T) {
	_, assert := test.Describe(t)

	tests := []struct {
		ts    []int
		start int
		len   int
		score int
	}{
		{[]int{-3, -5, -4, 0, 0, 2, 5, 6, 7}, 0, 3, 3},
		{[]int{-3, -5, -4, 0, 0, 2, 5, 6, 7, -3}, 9, 4, 3},
		{[]int{-191, -95, -191, -191, -191, -191, -191, -191, 0, 0, 0, -191}, 11, 9, 95},
		{[]int{0, 16, 0, 0, 0, 0, 0, 0, 143, 143, 143, 0}, 0, 0, 0},
	}
	for _, tt := range tests {
		co := continuousDark(tt.ts)
		assert.Equal(tt.start, co.start)
		assert.Equal(tt.len, co.length)
		assert.Equal(tt.score, co.score)
		assert.True(co.dark)
	}
}

func Test_continuousBright(t *testing.T) {
	_, assert := test.Describe(t)

	tests := []struct {
		ts    []int
		start int
		len   int
		score int
	}{
		{[]int{-3, -5, -4, 0, 0, 2, 5, 6, 7}, 5, 4, 2},
		{[]int{-3, -5, -4, 0, 0, 2, 5, 6, 7, -3}, 5, 4, 2},
	}
	for _, tt := range tests {
		co := continuousBright(tt.ts)
		assert.Equal(tt.start, co.start)
		assert.Equal(tt.len, co.length)
		assert.Equal(tt.score, co.score)
		assert.False(co.dark)
	}
}
