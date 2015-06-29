// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

// Problem3 is solution for finding the largest prime factor of the number 600851475143.
func Problem3() int {
	n := 600851475143
	lf := 1

	// Because 2 is the only even prime, we can treat it
	// separately and then increase 'f' with 2 on every step.
	if n%2 == 0 {
		lf = 2
		n /= 2
		for n%2 == 0 { // Divide out all 2 from 'n'.
			n /= 2
		}
	}

	// We check f*f <= n 'cause any number 'n' can have only one prime factor greater
	// than sqrt of 'n'. If 'n' is prime, so that prime factor is 'n' by its self.
	for f := 3; n > 1 && f*f <= n; f += 2 {
		if n%f == 0 {
			lf = f
			n /= f
			for n%f == 0 { // Divide out all the same factors from 'n'.
				n /= f
			}
		}
	}

	// If 'f' exceeds sqrt of 'n' we know the remaining number 'n' is
	// prime, because we've divided out all multiples 1..f from 'n'.
	if n != 1 {
		return n
	}
	return lf
}
