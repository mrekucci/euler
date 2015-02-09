// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

// Problem23 is solution for finding the the sum of all the positive
// integers which cannot be written as the sum of two abundant numbers.
func Problem23() int {
	const limit = 28123
	var abd []int
	// Were starting from 12 'cause it is the smallest abundant number.
	// The upper limit is  28123 'cause it can be shown that all greater
	// integers can be written as the sum of two abundant numbers.
	for i := 12; i <= limit; i++ {
		if isAbundant(i) {
			abd = append(abd, i)
		}
	}

	sums := make(map[int]bool)
	for i, a := range abd {
		for j := i; j < len(abd); j++ {
			s := a + abd[j]
			if s <= limit {
				sums[s] = true
			}
		}
	}

	sum := 0
	for i := 1; i <= limit; i++ {
		if !sums[i] {
			sum += i
		}
	}
	return sum
}

// isAbundant returns true if n is a abundant number.
// For n = 0 returns false, for n < 0 the result is not specified.
// A number is called abundant if sum of its divisors exceeds n.
func isAbundant(n int) bool {
	return PropDivSum(n) > n
}
