// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem004

// Solution for finding the largest palindrome
// made from the product of two 3-digit numbers.
func Solution() int {
	lp := 0
	// The palindrome can be written as: abccba which simplifies to:
	// 100000a + 10000b + 1000c + 100c + 10b + 1a.
	// We can simplify this to: 100001a + 10010b + 1100c.
	// Factoring out 11 we get: 11(9091a + 910b + 100c).
	// So the palindrome must be divisible by 11.
	// Because 11 is prime, at least one of the numbers must be divisible by 11.
	for a := 999; a > 99; a-- {
		b := 990 // First we assume that 'b' is divisible by 11.
		shift := 11
		if a%11 == 0 { // If 'a' is, then change 'b' to default.
			b = 999
			shift = 1
		}
		// To avoid checking numbers twice, e.g.: the number 69696 is checked
		// both, when a=132 and b=528, and when a=528 and b=132. To avoid this
		// double checks, we can assume that a ≤ b, and so halving the number
		// of calculations needed.
		for a <= b {
			n := a * b
			if n <= lp {
				break
			}
			if isPalindrome(n) {
				lp = n
			}
			b -= shift
		}
	}
	return lp
}

// reverse number n.
func reverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

// isPalindrome check if number n is a palindrome.
func isPalindrome(n int) bool {
	return n == reverse(n)
}
