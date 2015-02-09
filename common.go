// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "math"

// Integer limit values.
const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

// IsPrime returns true if number n is a prime number.
func IsPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n < 4:
		return true
	case n%2 == 0:
		return false
	case n < 9:
		return true
	case n%3 == 0:
		return false
	}

	// Any number 'n' can have only one prime factor greater than sqrt of 'n'.
	// And if 'n' is prime, so that prime factor is 'n' by its self.
	l := int(math.Sqrt(float64(n)))
	// Here we start with 5 because all primes greater than 3 can be written
	// in the form: (6k-/+1). So for: k=1 is (6k-1) equal 5.
	// Prime for next k-th prime is computed by: p+=6 'cause we started with (6k-1) for k=1.
	// WARNING: the series (6k-/+1) contains all primes but
	// not all numbers in this series are prime (e.g: for k=4 (6k+1) is not)!
	for p := 5; p <= l; p += 6 {
		if n%p == 0 { // We check here for (6k-1).
			return false
		}
		if n%(p+2) == 0 { // We check here for (6k+1).
			return false
		}
	}
	return true
}

// Mul returns result of a * b, and true when multiplication
// doesn't overflow, otherwise return 0 and false.
func Mul(a, b int) (int, bool) {
	switch {
	case a == 0 || b == 0:
		return 0, true
	case a == 1:
		return b, true
	case b == 1:
		return a, true
	// Next check involve the following cases:
	// If a=MinInt then a*b overflows for every -1<b>1.
	// If b=MinInt then b*a overflows for every -1<a>1.
	// If a=MinInt and b=-1 then c=a*b overflows with result c=MinInt, this causes
	// the division c/b to overflow (c/b=MinInt) and thus c/b=a which isn't really true.
	case a == MinInt || b == MinInt:
		return 0, false
	default:
		c := a * b
		if c/b != a {
			return 0, false
		}
		return c, true
	}
}

// MulInts returns a number which is multiplication of integers in slice s.
// A 0 and false is returned when multiplication overflows integer limits.
func MulInts(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, true
	}
	res := 1
	for _, n := range s {
		m, ok := Mul(res, n)
		if !ok {
			return 0, false
		}
		res = m
	}
	return res, true
}

// PropDivSum returns sum of proper divisors of positive integer n.
// For n = 0 returns 0, for n < 0 the result is not specified.
// Proper divisors are numbers less than n which divide evenly into n.
func PropDivSum(n int) int {
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
