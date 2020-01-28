package interp

// Algo is the base interface for interpolation algorithms
type Algo interface {
	// Value interpolates a single value
	Value(x float64) float64

	// Values interpolates a whole slice of values.  It will panic if
	// the slice is not sorted in ascending order.
	Values(dst, xs []float64) []float64

	// Fit refits the underlying interpolation
	// Needed if you change the underlying data
	Fit() error

	// SetRawData updates the underlying raw data and does
	// a refit
	SetRawData(xs, ys []float64) error

	// RawData returns the underlying raw data
	RawData() ([]float64, []float64)
}

// NumInputs returns the number of underlying points in an Algo
func NumInputs(a Algo) int {
	x, _ := a.RawData()
	return len(x)
}
