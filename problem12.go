// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

// Problem12 is solution for finding the value of the first
// triangle number to have over five hundred divisors.
func Problem12() int {
	tn := 0
	for i := 1; i <= MaxInt; i++ {
		tn += i
		dc := 2 // We start with 2 because every number is divisible by 1 and its self.
		// We go up to f*f <= tn 'cause any number 'tn' can have only one prime factor greater
		// than sqrt of 'tn'. If 'tn' is prime, so that prime factor is 'tn' by its self.
		//
		// We also check if f*f <= tn because if 'tn' is composite then
		// it consist from a*b where if a=b then a*a<=tn.
		for f := 2; f*f <= tn; f++ {
			if tn%f == 0 {
				dc++
				// We check up to 250 because the composite a*b is a reflection
				// of b*a after sqrt of 'tn', so we divided 500/2.
				if dc >= 250 {
					return tn
				}
			}
		}
	}

	panic("unreachable statement") // This case should never happen.
}
