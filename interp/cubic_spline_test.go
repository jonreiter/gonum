package interp_test

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/floats"

	"gonum.org/v1/gonum/interp"
)

const cubicsplinetol = 1e-5

type CubicSplineTestCase struct {
	xs        []float64
	ys        []float64
	testX     []float64
	expectedY []float64
}

func TestCubicSpline(t *testing.T) {
	var ltc CubicSplineTestCase
	ltc.xs = []float64{0, 2.5, 5, 7.5, 10}
	ltc.ys = []float64{0, 2.5, 5, 7.5, 10}
	ltc.testX = []float64{-1, 3, 6, 8, 11}
	ltc.expectedY = []float64{-1, 3, 6, 8, 11}

	l, err := interp.NewCubicSpline(ltc.xs, ltc.ys)
	if err != nil {
		t.Error("got back error building new linear")
	}

	for i, x := range ltc.testX {
		result := l.Value(x)
		expected := ltc.expectedY[i]
		if !floats.EqualWithinAbs(result, expected, cubicsplinetol) {
			t.Error("linear interp mismatch")
		}
	}
	results := l.Values(ltc.testX, nil)
	for i, v := range ltc.expectedY {
		if !floats.EqualWithinAbs(v, results[i], cubicsplinetol) {
			fmt.Println(i, " : ", v, " vs ", results[i])
			t.Error("values mismatch")
		}
	}
	dst := make([]float64, len(ltc.expectedY))
	results2 := l.Values(ltc.testX, dst)
	for i, v := range ltc.expectedY {
		if !floats.EqualWithinAbs(v, results2[i], cubicsplinetol) {
			t.Error("values mismatch")
		}
		if !floats.EqualWithinAbs(v, dst[i], cubicsplinetol) {
			t.Error("values mismatch")
		}
	}
}
