package interp

import (
	"errors"
	"sort"
)

// Error represents interp handling errors.
type Error string

func (err Error) Error() string { return string(err) }

var (
	// ErrNotSorted indicates unsorted slices were passes where sorted slices were expected
	ErrNotSorted = errors.New("interp: entries not sorted")
	// ErrLengthMismatch indicates slices of different lengths were passes where the same length was expected
	ErrLengthMismatch = errors.New("interp: slice length mismatch")
)

func panicIfNotSorted(x []float64) {
	if !sort.Float64sAreSorted(x) {
		panic(ErrNotSorted)
	}
}

func panicIfDifferentLengths(x, y []float64) {
	if len(x) != len(y) {
		panic(ErrLengthMismatch)
	}
}
