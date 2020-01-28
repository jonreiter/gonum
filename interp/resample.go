package interp

import (
	"gonum.org/v1/gonum/floats"
)

// Resample removes all underlying points where check returns true
// and replaces them with values taken from the passed-in
// interpolation Algo.
func Resample(r Algo, exclude func(float64) bool) error {
	rawX, rawY := r.RawData()
	inds, err := floats.Find(nil, func(f float64) bool {
		return !exclude(f)
	}, rawY, -1)
	if err != nil {
		return err
	}
	newx := make([]float64, len(inds))
	newy := make([]float64, len(inds))
	for i := range newx {
		newx[i] = rawX[inds[i]]
		newy[i] = rawY[inds[i]]
	}
	r.SetRawData(newx, newy)

	for i, v := range rawX {
		resampleThis := true
		for _, j := range inds {
			if i == j {
				resampleThis = false
			}
		}
		if resampleThis {
			rawY[i] = r.Value(v)
		}
	}
	r.SetRawData(rawX, rawY)
	return nil
}
