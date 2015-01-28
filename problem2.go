// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

// Problem2 is solution for finding the sum of the even-valued terms
// in the Fibonacci sequence whose values do not exceed four million.
func Problem2() int {
	sum := 0
	// Every 3rd number is even, so:
	// 1, 1, 2, 3, 5, 8, ...
	// a, b, c, a, b, c, ...     where c = a + b
	a := 1
	b := 1
	c := a + b
	for c <= 4e6 {
		sum += c
		a = b + c
		b = c + a
		c = a + b
	}
	return sum
}
