// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"math"
	"testing"
)

// AssertEquals fails by Errorf if res != want.
// It helps reduce boilerplate code in tests.
func AssertEquals(t *testing.T, res int, want int) {
	if res != want {
		t.Errorf("Result = %d; want %d", res, want)
	}
}

// IsPrime returns true if number n is a prime number.
func IsPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n < 4:
		return true
	case n%2 == 0:
		return false
	case n < 9:
		return true
	case n%3 == 0:
		return false
	}
	// Any number 'n' can have only one prime factor greater than sqrt of 'n'.
	// And if 'n' is prime, so that prime factor is 'n' by its self.
	l := int(math.Sqrt(float64(n)))
	// Here we start with 5 because all primes greater than 3 can be written
	// in the form: (6k-/+1). So for: k=1 is (6k-1) equal 5.
	// Prime for next k-th prime is computed by: p+=6 'cause we started with (6k-1) for k=1.
	// WARNING: the series (6k-/+1) contains all primes but
	// not all numbers in this series are prime (e.g: for k=4 (6k+1) is not)!
	for p := 5; p <= l; p += 6 {
		if n%p == 0 { // We check here for (6k-1).
			return false
		}
		if n%(p+2) == 0 { // We check here for (6k+1).
			return false
		}
	}
	return true
}
