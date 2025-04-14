// v0.1.0

package imgpath

import (
	"github.com/wunderbarb/test"
	"testing"
)

func TestCountContinuousBitsAtOne(t *testing.T) {
	require, _ := test.Describe(t)

	tests := []struct {
		n     uint32
		start int
		len   int
	}{
		{0b011100110, 5, 3},
		{0b111101100, 5, 4},
		{0b111101101, 5, 5},
	}
	for _, tt := range tests {
		s, l := CountContinuousBitsAtOne(tt.n, 9)
		require.Equal(tt.len, l)
		require.Equal(tt.start, s)
	}
}
