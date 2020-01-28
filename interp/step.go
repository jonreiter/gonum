package interp

import (
	"sort"
)

// Step does step interpolation
type Step struct {
	x []float64
	y []float64
}

// NewStep builds and fits a new linear interpolator.
func NewStep(x, y []float64) (*Step, error) {
	var l Step
	l.x = x
	l.y = y
	err := l.Fit()
	if err != nil {
		return nil, err
	}
	return &l, nil
}

// Fit refits the undelying slices.
func (l *Step) Fit() error {
	panicIfDifferentLengths(l.x, l.y)
	return nil
}

// Value returns the linearly interpolated value
func (l *Step) Value(x float64) float64 {
	val, _ := l.stepWithIndex(x, 0)
	return val
}

// Values interpolates a sorted slice of values
func (l *Step) Values(dst, xs []float64) []float64 {
	if dst == nil {
		dst = make([]float64, len(xs))
	}
	if len(xs) != len(dst) {
		panic(ErrLengthMismatch)
	}
	startIdx := 0
	for i, x := range xs {
		y, idx := l.stepWithIndex(x, startIdx)
		startIdx = idx
		dst[i] = y
	}
	return dst
}

func (l *Step) stepWithIndex(x float64, startIdx int) (value float64, index int) {
	xs := l.x[startIdx:]
	ys := l.y[startIdx:]
	idx := sort.SearchFloat64s(xs, x)
	if idx == 0 {
		return ys[0], idx + startIdx
	}
	if idx == NumInputs(l)-startIdx {
		return ys[NumInputs(l)-startIdx-1], idx + startIdx
	}
	return ys[idx-1], idx + startIdx
}

// RawData returns the raw underling points.
func (l *Step) RawData() ([]float64, []float64) {
	return l.x, l.y
}

// SetRawData updates the raw underling points.
func (l *Step) SetRawData(xs, ys []float64) error {
	l.x = xs
	l.y = ys
	return l.Fit()
}
