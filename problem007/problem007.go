// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem007

import "github.com/mrekucci/euler/util"

// Solution for finding what is the 10 001st prime number.
func Solution() int {
	count := 1 // Already include number 2, because we start from 3.
	pc := 1
	for count < 10001 {
		// All primes except 2 are odd. Odd number is defined by: (2k + 1).
		// Starting by odd number and moving by 2, we will iterate through all odd numbers.
		pc += 2
		if util.IsPrime(pc) {
			count++
		}
	}
	return pc
}
