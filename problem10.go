// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

// Problem10 is solution for finding the sum of all the primes below two million.
func Problem10() int {
	sum := 2                      // We assigned 2 because it's the first prime.
	for i := 3; i < 2e6; i += 2 { // Starting by 3 and moving be 2 we are skipping every even number.
		if IsPrime(i) {
			sum += i
		}
	}
	return sum
}
