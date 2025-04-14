package imgpath

import (
	"image"
	"math"
)

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
	ip.index = 0
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
	ip.SetCenter(cai.X, cai.Y)
	vc := ip.grayCenter()
	if vc <= cai.T {
		return
	}
	ip.index = 0
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
	ip.SetCenter(cai.X, cai.Y)
	vc := ip.grayCenter()
	if vc >= 0xff-cai.T {
		return
	}
	p := vc + cai.T
	score = 0xFF
	l := 0
	ok = false
	ip.Until(func(v uint8, index int) bool {
		if ip.Next() >= p {
			l++
			score = min(ip.Next()-p, score)
			if l < cai.Length {
				return true
			}
			ok = true
			return false
		}
		l = 0
		score = 0xFF
		return !ip.Cycled()
	})
	return
}

// ContinuousDarkerAtLeast checks whether the path as at least `Length` consecutive pixels darker with the threshold
// `T`.  It returns the score (the higher, the better) and the momentum `angle`.
func (ip ImagePath) ContinuousDarkerAtLeast(cai ContinuousInput) (score uint8, ok bool) {
	ip.SetCenter(cai.X, cai.Y)
	vc := ip.grayCenter()
	p := cai.T - vc
	score = 0xFF
	l := 0
	ip.Until(func(v uint8, index int) bool {
		if ip.Next() <= p {
			l++
			score = min(p-ip.Next(), score)
			if l >= cai.Length {
				return true
			}
		}
		l = 0
		return !ip.Cycled()
	})
	return
}

func (ip ImagePath) fast4Cascade(ci ContinuousInput) (score uint8, angle float64, dark bool, ok bool) {
	score, angle, ok = ip.ContinuousBrighter(ci)
	if ok {
		dark = false
		return
	}
	score, angle, ok = ip.ContinuousDarker(ci)
	if ok {
		dark = true
		return
	}
	return
}

// CascadedFast calculates the cascaded fast of point x, y.
// https://ealdea.github.io/VisionSETI/pdfs/cascaded-fast.pdf
func CascadedFast(img *image.Gray, x, y int, t uint8) (score uint8, angle float64, ok bool) {
	const (
		alpha = math.Pi / 10
		beta  = math.Pi / 8
	)
	C3.SetImage(img)
	s1, a1, dark1, ok1 := C3.fast4Cascade(ContinuousInput{
		X:      x,
		Y:      y,
		T:      t,
		Length: 6,
	})
	if !ok1 {
		ok = false
		return
	}
	C4.SetImage(img)
	s2, a2, dark2, ok2 := C4.fast4Cascade(ContinuousInput{
		X:      x,
		Y:      y,
		T:      t,
		Length: 9,
	})
	if !ok2 || dark1 != dark2 {
		ok = false
		return
	}
	C5.SetImage(img)
	s3, a3, dark3, ok3 := C5.fast4Cascade(ContinuousInput{
		X:      x,
		Y:      y,
		T:      t,
		Length: 11,
	})
	if !ok3 || dark1 != dark3 {
		ok = false
		return
	}
	ok = math.Abs(a1-a3) <= alpha && math.Abs(a1-a2) <= beta
	if !ok {
		return
	}
	score = min(s1, s2, s3)
	angle = a3
	return
}
