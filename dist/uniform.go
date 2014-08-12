// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dist

import (
	"math"
	"math/rand"
)

type Uniform struct {
	Min float64
	Max float64
}

// CDF computes the value of the cumulative density function at x.
func (u Uniform) CDF(x float64) float64 {
	if x < u.Min {
		return 0
	}
	if x > u.Max {
		return 1
	}
	return (x - u.Min) / (u.Max - u.Min)
}

// Uniform doesn't have any of the DLogProbD? because the derivative is 0 everywhere
// except where it's undefined

// Entropy returns the entropy of the distribution.
func (u Uniform) Entropy() float64 {
	return math.Log(u.Max - u.Min)
}

// ExKurtosis returns the excess kurtosis of the distribution.
func (Uniform) ExKurtosis() float64 {
	return -6.0 / 5.0
}

// Uniform doesn't have Fit because it's a bad idea to fit a uniform from data.

// LogProb computes the natural logarithm of the value of the probability density function at x.
func (u Uniform) LogProb(x float64) float64 {
	return -math.Log(u.Max - u.Min)
}

// Mean returns the mean of the probability distribution.
func (u Uniform) Mean() float64 {
	return (u.Max - u.Min) / 2
}

// Median returns the median of the probability distribution.
func (u Uniform) Median() float64 {
	return (u.Max - u.Min) / 2
}

// Uniform doesn't have a mode because it's any value in the distribution

// NumParameters returns the number of parameters in the distribution.
func (Uniform) NumParameters() int {
	return 2
}

// Parameters gets the parameters of the distribution. Panics if the length of
// the input slice is non-zero and not equal to the number of parameters.
// This returns a slice with length 2 containing the minimum and maximum of
// the distribution.
func (u Uniform) Parameters(s []float64) []float64 {
	nParam := u.NumParameters()
	if s == nil {
		s = make([]float64, nParam)
	}
	if len(s) != nParam {
		panic("uniform: improper parameter length")
	}
	s[0] = u.Min
	s[1] = u.Max
	return s
}

// Prob computes the value of the probability density function at x.
func (u Uniform) Prob(x float64) float64 {
	return 1 / (u.Max - u.Min)
}

// Quantile returns the inverse of the cumulative probability distribution.
func (u Uniform) Quantile(p float64) float64 {
	if p < 0 || p > 1 {
		panic("dist: percentile out of bounds")
	}
	return p*(u.Max-u.Min) + u.Min
}

// Rand returns a random sample drawn from the distribution.
func (u Uniform) Rand() float64 {
	return rand.Float64()*(u.Max-u.Min) + u.Min
}

// SetParameters gets the parameters of the distribution. Panics if the length of
// the input slice is not equal to the number of parameters.
// This sets the minimum to the first element and the maximum to the second
// element.
func (u *Uniform) SetParameters(s []float64) {
	if len(s) != u.NumParameters() {
		panic("uniform: incorrect number of parameters to set")
	}
	u.Min = s[0]
	u.Max = s[1]
}

// Skewness returns the skewness of the distribution.
func (Uniform) Skewness() float64 {
	return 0
}

// StdDev returns the standard deviation of the probability distribution.
func (u Uniform) StdDev() float64 {
	return math.Sqrt(u.Variance())
}

// Survival returns the survival function (complementary CDF) at x.
func (u Uniform) Survival(x float64) float64 {
	if x < u.Min {
		return 1
	}
	if x > u.Max {
		return 0
	}
	return (u.Max - x) / (u.Max - u.Min)
}

// Variance returns the variance of the probability distribution.
func (u Uniform) Variance() float64 {
	return 1.0 / 12.0 * (u.Max - u.Min) * (u.Max - u.Min)
}
