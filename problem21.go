// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

// Problem21 is solution for finding the sum of all the amicable numbers under 10000.
func Problem21() int {
	sum := 0
	for n := 2; n < 10e3; n++ { // We are starting by 2 'cause 1 has no divisors at all.
		if isAmicable(n) {
			sum += n
		}
	}
	return sum
}

// isAmicable decide if a is a amicable number.
// The number is amicable if next statement is true:
// Let d(n) be defined as the sum of proper divisors of
// n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a != b, then a and b
// are an amicable pair and each of a and b are called
// amicable numbers.
func isAmicable(a int) bool {
	b := PropDivSum(a)
	return a != b && a == PropDivSum(b)
}
