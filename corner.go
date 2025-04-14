package imgpath

import "math"

// ContinuousInput is the input to methods ContinuousBrighter and ContinuousDarker.
type ContinuousInput struct {
	// X is the horizontal coordinate of the center.
	X int
	// Y is the vertical coordinate of the center.
	Y      int
	T      uint8
	Length int
}

// ContinuousBrighter checks whether the path as at least `Length` consecutive pixels brighter with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousBrighter(cai ContinuousInput) (score uint8, angle float64, ok bool) {
	vc := ip.grayCenter()
	if vc >= cai.T {
		return
	}
	p := cai.T - vc
	var nn uint32
	score = 0xFF
	ip.All(func(v uint8, index int) {
		nn <<= 1
		if ip.Next() >= p {
			nn++
			score = min(ip.Next()-p, score)
		}
	})
	start, l := CountContinuousBitsAtOne(nn, ip.Len())
	if l < cai.Length {
		return
	}
	ok = true
	a := (start + l + 1) % ip.Len()
	angle = math.Pi * float64(a) / float64(ip.Len())
	return
}

// ContinuousDarker checks whether the path as at least `Length` consecutive pixels darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousDarker(cai ContinuousInput) (score uint8, angle float64, ok bool) {
	vc := ip.grayCenter()
	if vc <= cai.T {
		return
	}
	p := vc - cai.T
	var nn uint32
	score = 0xFF
	ip.All(func(v uint8, index int) {
		nn <<= 1
		if ip.Next() <= p {
			nn++
			score = min(p-ip.Next(), score)
		}
	})
	start, l := CountContinuousBitsAtOne(nn, ip.Len())
	if l < cai.Length {
		return
	}
	ok = true
	a := (start + l + 1) % ip.Len()
	angle = math.Pi * float64(a) / float64(ip.Len())
	return
}

// ContinuousBrighterAtLeast checks whether the path as at least `Length` consecutive pixels brighter with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousBrighterAtLeast(cai ContinuousInput) (score uint8, ok bool) {
	score, _, ok = ip.ContinuousBrighter(cai)
	return
}

// ContinuousDarkerAtLeast checks whether the path as at least `Length` consecutive pixels darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousDarkerAtLeast(cai ContinuousInput) (score uint8, ok bool) {
	score, _, ok = ip.ContinuousBrighter(cai)
	return
}
