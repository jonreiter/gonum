package interp

import "fmt"

// ExampleLinear shows how to use linear interp
func ExampleLinear() {
	xs := []float64{0, 2.5, 5, 7.5, 10}
	ys := []float64{2, 9.5, 17, 24.5, 32}
	l, err := NewLinear(xs, ys)

	if err != nil {
		panic(err)
	}

	interpX := 0.23
	interpY := l.Value(0.23)
	fmt.Println("for ", interpX, "interped value is", interpY)
}

// ExampleSpline shows how to use cubic spline interp
func ExampleSpline() {
	xs := []float64{0, 2.5, 5, 7.5, 10}
	ys := []float64{2, 9.5, 17, 24.5, 32}
	l, err := NewCubicSpline(xs, ys)

	if err != nil {
		panic(err)
	}

	interpX := 0.23
	interpY := l.Value(0.23)
	fmt.Println("for ", interpX, "interped value is", interpY)

	interpXSlice := []float64{2.2, 3.3, 4.4}
	interpYSlice := l.Values(nil, interpXSlice)
	fmt.Println("for ", interpXSlice, "interped value is", interpYSlice)
}
