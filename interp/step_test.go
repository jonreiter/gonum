package interp_test

import (
	"testing"

	"gonum.org/v1/gonum/floats"

	"gonum.org/v1/gonum/interp"
)

const steptol = 1e-5

type StepTestCase struct {
	xs        []float64
	ys        []float64
	testX     []float64
	expectedY []float64
}

func TestStep(t *testing.T) {
	var ltc StepTestCase
	ltc.xs = []float64{0, 2.5, 5, 7.5, 10}
	ltc.ys = []float64{2, 9.5, 17, 24.5, 32}
	ltc.testX = []float64{-1, 3, 6, 8, 11}
	ltc.expectedY = []float64{2, 9.5, 17, 24.5, 32}

	l, err := interp.NewStep(ltc.xs, ltc.ys)
	if err != nil {
		t.Error("got back error building new step")
	}
	for i, x := range ltc.testX {
		result := l.Value(x)
		expected := ltc.expectedY[i]
		if !floats.EqualWithinAbs(result, expected, steptol) {
			t.Error("step interp mismatch")
		}
	}
	results := l.Values(nil, ltc.testX)
	for i, v := range ltc.expectedY {
		if !floats.EqualWithinAbs(v, results[i], steptol) {
			t.Error("values mismatch")
		}
	}
	dst := make([]float64, len(ltc.expectedY))
	results2 := l.Values(dst, ltc.testX)
	for i, v := range ltc.expectedY {
		if !floats.EqualWithinAbs(v, results2[i], steptol) {
			t.Error("values mismatch")
		}
		if !floats.EqualWithinAbs(v, dst[i], steptol) {
			t.Error("values mismatch")
		}
	}

	ltc.ys[4] = 12
	err = l.Fit()
	if err != nil {
		t.Error("got back error refitting")
	}
	res := l.Value(11)
	if res != 12 {
		t.Error("refit error")
	}

}
