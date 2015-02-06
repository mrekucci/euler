// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "math"

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
	b := propDivSum(a)
	return a != b && a == propDivSum(b)
}

// propDivSum returns sum of proper divisors of positive integer n.
// For n = 0 returns 0, and if n is negative integer, the result is not specified.
// Proper divisors are numbers less than n which divide evenly into n.
func propDivSum(n int) int {
	// Let σ(n) be the sum of divisors of the natural number, n.
	// For any prime, p: σ(p) = p + 1, as the only divisors would be 1 and p.
	// σ(p^a) = (p^(a+1) − 1) / (p − 1).
	// σ(x * y * ...) = σ(x) * σ(y) * ..., where x, y, ..., are relatively prime.

	if n == 0 || n == 1 {
		return 0
	}

	cn := n
	sum := 1

	// Because 2 is the only even prime, we can treat it
	// separately and then increase 'p' with 2 on every step.
	if n%2 == 0 {
		pa := 2 * 2
		n /= 2
		for n%2 == 0 {
			pa *= 2
			n /= 2
		}
		sum *= (pa - 1) // (p^(a+1) - 1) / (p - 1) where p=2.
	}

	mf := int(math.Sqrt(float64(n)))
	// Any number 'n' can have only one prime factor greater than sqrt of 'n'.
	// If 'n' is prime, so that prime factor is 'n' by its self.
	for p := 3; n > 1 && p <= mf; p += 2 {
		if n%p == 0 {
			pa := p * p
			n /= p
			for n%p == 0 {
				pa *= p
				n /= p
			}
			sum *= ((pa - 1) / (p - 1)) // σ(a) * σ(b) * ...
		}
	}

	if n != 1 { // If one prime factor greater then sqrt of 'n' remains, add it to the sum.
		sum *= (n + 1)
	}

	return sum - cn // The sum of the proper divisors of 'n' is the 'sum' of its divisors minus 'n' itself.
}
