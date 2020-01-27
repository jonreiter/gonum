package interp

import (
	"sort"
)

// Linear does basic piecewise linear interpolation
type Linear struct {
	x []float64
	y []float64
}

// NewLinear builds and fits a new linear interpolator.
func NewLinear(x, y []float64) (*Linear, error) {
	var l Linear
	l.x = x
	l.y = y
	err := l.Fit()
	if err != nil {
		return nil, err
	}
	return &l, nil
}

// Fit refits the undelying slices.
func (l *Linear) Fit() error {
	panicIfDifferentLengths(l.x, l.y)
	return nil
}

// NumInputs returns the number of input points
func (l *Linear) NumInputs() int {
	return len(l.x)
}

// Value returns the linearly interpolated value
func (l *Linear) Value(x float64) float64 {
	val, _ := l.linearPiecewiseWithIndex(x, 0)
	return val
}

// Values interpolates a sorted slice of values
func (l *Linear) Values(xs, dst []float64) []float64 {
	if dst == nil {
		dst = make([]float64, len(xs))
	}
	if len(xs) != len(dst) {
		panic(ErrLengthMismatch)
	}
	startIdx := 0
	for i, x := range xs {
		y, idx := l.linearPiecewiseWithIndex(x, startIdx)
		startIdx = idx
		dst[i] = y
	}
	return dst
}

func (l *Linear) linearPiecewiseWithIndex(x float64, startIdx int) (value float64, index int) {
	xs := l.x[startIdx:]
	ys := l.y[startIdx:]
	idx := sort.SearchFloat64s(xs, x)
	if idx == 0 {
		return ys[0], idx + startIdx
	}
	if idx == l.NumInputs()-startIdx {
		return ys[l.NumInputs()-startIdx-1], idx + startIdx
	}
	//	fmt.Println("got idx: ", idx)
	return linearPiecewise(x, xs[idx-1], xs[idx], ys[idx-1], ys[idx]), idx + startIdx
}

func linearPiecewise(x, x0, x1, y0, y1 float64) float64 {
	dx := x1 - x0
	dy := y1 - y0
	dRatio := dy / dx
	return y0 + (x-x0)*dRatio
}
