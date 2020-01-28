package interp

import (
	"math"
	"sort"
)

type spline struct {
	a float64
	b float64
	c float64
	d float64
	x float64
}

func (s *spline) value(x float64) float64 {
	return s.a + s.b*(x-s.x) + s.c*math.Pow(x-s.x, 2) + s.d*math.Pow(x-s.x, 3)
}

// CubicSpline interpolation
type CubicSpline struct {
	x       []float64
	y       []float64
	splines []spline
}

// NewCubicSpline builds and fits a new cubic spline interpolator.
func NewCubicSpline(x, y []float64) (*CubicSpline, error) {
	var l CubicSpline
	l.x = x
	l.y = y
	err := l.Fit()
	if err != nil {
		return nil, err
	}
	return &l, nil
}

// Fit refits the undelying slices.
func (s *CubicSpline) Fit() error {
	panicIfDifferentLengths(s.x, s.y)
	n := NumInputs(s) - 1

	a := make([]float64, n+1)
	for i, v := range s.y {
		a[i] = v
	}
	b := make([]float64, n)
	d := make([]float64, n)

	h := make([]float64, n)
	for i := 0; i < n; i++ {
		h[i] = s.x[i+1] - s.x[i]
	}

	alpha := make([]float64, n)
	for i := 1; i < n; i++ {
		alpha[i] = (3/h[i])*(a[i+1]-a[i]) - (3/h[i-1])*(a[i]-a[i-1])
	}

	c := make([]float64, n+1)
	l := make([]float64, n+1)
	mu := make([]float64, n+1)
	z := make([]float64, n+1)

	l[0] = 1
	mu[0] = 0
	z[0] = 0

	for i := 1; i < n; i++ {
		l[i] = 2*(s.x[i+1]-s.x[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}

	l[n] = 1
	z[n] = 0
	c[n] = 0

	for j := n - 1; j > -1; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
		b[j] = (a[j+1]-a[j])/(h[j]) - (h[j]*(c[j+1]+2*c[j]))/(3)
		d[j] = (c[j+1] - c[j]) / (3 * h[j])
	}

	s.splines = make([]spline, n)
	for i := range s.splines {
		s.splines[i].a = a[i]
		s.splines[i].b = b[i]
		s.splines[i].c = c[i]
		s.splines[i].d = d[i]
		s.splines[i].x = s.x[i]
	}

	return nil
}

// Value returns the linearly interpolated value
func (s *CubicSpline) Value(x float64) float64 {
	val, _ := s.stepWithIndex(x, 0)
	return val
}

// Values interpolates a sorted slice of values
func (s *CubicSpline) Values(dst, xs []float64) []float64 {
	if dst == nil {
		dst = make([]float64, len(xs))
	}
	if len(xs) != len(dst) {
		panic(ErrLengthMismatch)
	}
	startIdx := 0
	for i, x := range xs {
		y, idx := s.stepWithIndex(x, startIdx)
		startIdx = idx
		dst[i] = y
	}
	return dst
}

func (s *CubicSpline) stepWithIndex(x float64, startIdx int) (value float64, index int) {
	xs := s.x[startIdx:]
	idx := sort.SearchFloat64s(xs, x)
	if idx == 0 {
		return s.splines[0].value(x), idx
	}
	if idx > NumInputs(s)-2 {
		return s.splines[NumInputs(s)-2].value(x), idx
	}
	return s.splines[idx-1].value(x), idx
}

// RawData returns the raw underlying points
func (s *CubicSpline) RawData() ([]float64, []float64) {
	return s.x, s.y
}

// SetRawData sets the raw underlying points
func (s *CubicSpline) SetRawData(x, y []float64) error {
	s.x = x
	s.y = y
	return s.Fit()
}
