// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "math"

// Problem5 is solution for finding the smallest positive number that
// is evenly divisible by all of the numbers from 1 to 20.
func Problem5() int {
	// The first solution:
	//
	// 	It is like finding the least common multiple of set 1-20:
	// 	1, 2, 3, 4=2*2, 5, 6=2*3, 7, 8=2*2*2, 9=3*3, 10=2*5, 11, 12=2*2*3,
	// 	13, 14=2*7, 15=3*5, 16=2*2*2*2, 17, 18=2*3*3, 19, 20=2*2*5
	//	Next multiple include all maximum occurrences of prime factors in numbers 2 to 20
	// 	return 2 * 2 * 2 * 2 * 3 * 3 * 5 * 7 * 11 * 13 * 17 * 19
	//

	// The second solution, generalized:
	var n float64 = 1
	var k float64 = 20                              // The upper limit of the range.
	primes := []float64{2, 3, 5, 7, 11, 13, 17, 19} // List of all primes from range 1-20.
	// 'l' represents limit where we only need to evaluate exponent, after this limit
	// exponent is equal 1. This solution is optimization only for case 1-20!
	l := math.Sqrt(k)
	for _, p := range primes {
		var exp float64 = 1
		// Note that exp=1 for p^2>k; this solution is optimization only for case 1-20!
		if p <= l {
			// p^exp=k; after logarithmic both sides we get: exp*log(p)=log(k).
			// From which we get: exp=log(k)/log(prime); because the p^exp<=k
			// the exponent must be floor to integer.
			exp = math.Floor(math.Log(k) / math.Log(p))
		}
		// The greatest perfect power of 'p' that is less than or equal to 'k'.
		// No other numbers in 'primes' can contain more than 'exp' times repeated factors of 'p'.
		n *= math.Pow(p, exp)
	}
	return int(n)
}
