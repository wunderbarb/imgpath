// v0.2.0
// Author: wunderbarb
// (c), Apr 2025

package imgpath

// ContinuousInput is the input to methods ContinuousBrighter and ContinuousDarker.
type ContinuousInput struct {
	// X is the horizontal coordinate of the center.
	X int
	// Y is the vertical coordinate of the center.
	Y int
	// T is the threshold.
	T      uint8
	Length int
}

// ContinuousBrighter checks whether the path as at least `Length` consecutive pixels brighter with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousBrighter(cai ContinuousInput) (uint8, float64, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousBright)
	if co.length < cai.Length {
		return 0, 0.0, false
	}
	return uint8(co.score), co.angle, true
}

// ContinuousBrighterExact checks whether the path as exactly `Length` consecutive pixels brighter with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.  It returns the first encountered
// successful sequence.
func (ip ImagePath) ContinuousBrighterExact(cai ContinuousInput) (uint8, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousBright)
	if co.length != cai.Length {
		return 0, false
	}
	return uint8(co.score), true
}

// ContinuousDarker checks whether the path as at least `Length` consecutive pixels darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousDarker(cai ContinuousInput) (uint8, float64, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousDark)
	if co.length < cai.Length {
		return 0, 0.0, false
	}
	return uint8(co.score), co.angle, true
}

// ContinuousDarkerExact checks whether the path as exactly `Length` consecutive pixels darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousDarkerExact(cai ContinuousInput) (uint8, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousDark)
	if co.length != cai.Length {
		return 0, false
	}
	return uint8(co.score), true
}

func (ip *ImagePath) process(x, y int, t uint8, fn func([]int) continuousOutput) continuousOutput {
	ip.SetCenter(x, y)
	ip.Reset()
	ts := ip.Diff()
	tsp := thresholdDiff(ts, t)
	return fn(tsp)
}

func thresholdDiff(ts []int, t uint8) []int {
	ts1 := make([]int, len(ts))
	for i, v := range ts {
		switch {
		case v > int(t):
			ts1[i] = v - int(t)
		case v < -int(t):
			ts1[i] = v + int(t)
		}
	}
	return ts1
}
