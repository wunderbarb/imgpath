// v0.1.0

package imgpath

import (
	"errors"
	"math"
)

type IndexRing struct {
	index  int
	length int
	cycled bool
}

// NewIndexRing creates a new IndexRing of size `l`.  `l` must be greater than one.
func NewIndexRing(l int) (*IndexRing, error) {
	if l <= 1 {
		return nil, errors.New("size should be at least 2")
	}
	return &IndexRing{length: l}, nil
}

func (ir IndexRing) Cycled() bool {
	return ir.cycled
}

// Next returns the current value of the index and points to the next one.
func (ir *IndexRing) Next() int {
	ii := ir.index
	ir.index++
	if ir.index >= ir.length {
		ir.index = 0
		ir.cycled = true
	}
	return ii
}

func (ir *IndexRing) Reset() {
	ir.index = 0
	ir.cycled = false
}

type continuousOutput struct {
	length int
	start  int
	angle  float64
	score  int
	dark   bool
}

func continuousBright(ts []int) continuousOutput {
	co := continuous(ts, func(in int) bool { return in > 0 })
	co.dark = false
	return co

}

func continuousDark(ts []int) continuousOutput {
	co := continuous(ts, func(in int) bool { return in < 0 })
	co.dark = true
	return co

}

func continuous(ts []int, fn func(in int) bool) continuousOutput {
	var co continuousOutput
	ir, err := NewIndexRing(len(ts))
	if err != nil {
		return co
	}
	var l, start int
	continuous := false
	score := 0xFF
	for {
		ii := ir.Next()
		if fn(ts[ii]) {
			l++
			score = min(score, abs(ts[ii]))
			if !continuous {
				continuous = true
				start = ii
			}
			continue
		}
		if continuous {
			continuous = false
			if l > co.length {
				co.length = l
				co.start = start
				co.score = score
				a := (3*start + l) % len(ts)
				co.angle = math.Pi * float64(a) / float64(len(ts))
			}
			l = 0
			score = 0xFF
		}
		if ir.Cycled() {
			return co
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
