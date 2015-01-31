// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

// Problem14 is solution for finding a starting number, under one
// million, which produces the longest chain by Collatz sequence.
func Problem14() int {
	sn := 0
	max := 0
	for i := 1; i < 1e6; i++ {
		n := collatz(i)
		count := 1
		for n != 1 {
			n = collatz(n)
			count++
		}
		if count > max {
			max = count
			sn = i
		}
	}
	return sn
}

// collatz returns:
// 	n/2    if n is even
// 	3n + 1 if n is odd
//
func collatz(n int) int {
	if n%2 == 0 {
		return n >> 1
	}
	return 3*n + 1
}
