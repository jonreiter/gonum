package interp

// Algo is the base interface for interpolation algorithms
type Algo interface {
	// Value interpolates a single value
	Value(x float64) float64

	// Values interpolates a whole slice of values.  It will panic if
	// the slice is not sorted in ascending order.
	Values(xs, dst []float64) []float64

	// Fit refits the underlying interpolation
	// Needed if you change the underlying data
	Fit() error

	// NumInputs returns the number of input points
	NumPoints() int
}
