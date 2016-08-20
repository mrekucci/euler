// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

// Problem3 is solution for finding the largest prime factor of the number 600851475143.
func Problem3() int {
	n, lf := 600851475143, 1

	// factorOut factors out all the same factors f from n if n%f==0, and sets lf to f.
	factorOut := func(f int) {
		if n%f == 0 {
			lf = f
			n /= f
			for n%f == 0 {
				n /= f
			}
		}
	}

	// Because 2 is the only even prime, we can treat it
	// separately and then increase f with 2 on every step.
	factorOut(2)

	// We check f*f <= n because any number n can have only one prime factor greater
	// than sqrt(n). If n is prime, so that prime factor is n by its self.
	for f := 3; n > 1 && f*f <= n; f += 2 {
		factorOut(f)
	}

	// If f exceeds sqrt(n) we know the remaining number n is a
	// prime, because we've factored out all multiples 1..f from n.
	if n != 1 {
		return n
	}
	return lf
}
