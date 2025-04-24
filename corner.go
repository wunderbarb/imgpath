// v0.4.0
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
	T uint8
	// Length defines the size of the continuous segment.
	Length int
}

type ContinuousOutput struct {
	// Length is the number of consecutive pixels.
	Length int
	// Score is the score of the path.
	Score uint8
	// Angle is the angle of the path in degree.
	Angle int
	// Darker is true if the path is darker.
	Darker bool
}

// AllBrighter checks whether all pixels in the path are brighter or equal to `threshold`.
func (ip ImagePath) AllBrighter(threshold uint8) bool {
	ip.Reset()
	for i := 0; i < len(ip.path); i++ {
		v := ip.Next()
		if v < threshold {
			return false
		}
	}
	return true
}

// AllDarker checks whether all pixels in the path are darker or equal to `threshold`.
func (ip ImagePath) AllDarker(threshold uint8) bool {
	ip.Reset()
	for i := 0; i < len(ip.path); i++ {
		if ip.Next() > threshold {
			return false
		}
	}
	return true
}

// Continuous checks whether the path has at least `Length` consecutive pixels brighter or darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.  It returns the first encountered
// successful sequence.  It returns the `Darker` flag to indicate whether the path is darker or brighter.
func (ip ImagePath) Continuous(cai ContinuousInput) (ContinuousOutput, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousBright)
	if co.length >= cai.Length {
		return ContinuousOutput{
			Length: co.length,
			Score:  uint8(co.score), // #nosec G115
			Angle:  co.angle,
			Darker: false,
		}, true
	}
	co = ip.process(cai.X, cai.Y, cai.T, continuousDark)
	if co.length >= cai.Length {
		return ContinuousOutput{
			Length: co.length,
			Score:  uint8(co.score), // #nosec G115
			Angle:  co.angle,
			Darker: true,
		}, true
	}
	return ContinuousOutput{}, false
}

// ContinuousBrighter checks whether the path as at least `Length` consecutive pixels brighter with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle` (in degrees).
func (ip ImagePath) ContinuousBrighter(cai ContinuousInput) (uint8, int, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousBright)
	if co.length < cai.Length {
		return 0, 0.0, false
	}
	return uint8(co.score), co.angle, true // #nosec G115
}

func (ip ImagePath) ContinuousBrighterThan(cai ContinuousInput) bool {
	ts := ip.Than(cai.T)
	co := continuousBright(ts)
	return co.length >= cai.Length
}

// ContinuousBrighterExact checks whether the path as exactly `Length` consecutive pixels brighter with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.  It returns the first encountered
// successful sequence.
func (ip ImagePath) ContinuousBrighterExact(cai ContinuousInput) (uint8, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousBright)
	if co.length != cai.Length {
		return 0, false
	}
	return uint8(co.score), true // #nosec G115
}

// ContinuousDarker checks whether the path as at least `Length` consecutive pixels darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle` (in degree).
func (ip ImagePath) ContinuousDarker(cai ContinuousInput) (uint8, int, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousDark)
	if co.length < cai.Length {
		return 0, 0.0, false
	}
	return uint8(co.score), co.angle, true // #nosec G115
}

// ContinuousDarkerExact checks whether the path as exactly `Length` consecutive pixels darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousDarkerExact(cai ContinuousInput) (uint8, bool) {
	co := ip.process(cai.X, cai.Y, cai.T, continuousDark)
	if co.length != cai.Length {
		return 0, false
	}
	return uint8(co.score), true //nolint:gosec
}

func (ip ImagePath) ContinuousDarkerThan(cai ContinuousInput) bool {
	ts := ip.Than(cai.T)
	co := continuousDark(ts)
	return co.length >= cai.Length
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
