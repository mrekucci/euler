// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

// Problem9 is solution for finding the product a*b*c of
// exactly one Pythagorean triplet for which a + b + c = 1000.
func Problem9() int {
	const sum = 1000
	var a, b, c int

	// sum/3 because a+b+c=sum and a<b<c => a+a+a < a+b+c => 3a<1000 => a<1000/3'.
	for a = 1; a < sum/3; a++ {
		// b=a+1 because a<b; sum/2 because b<c => b+b < b+c = 1000-a => 2b<1000-a => b<(1000-a)/2.
		for b = a + 1; b < (sum-a)/2; b++ {
			c = 1000 - a - b
			if c*c == a*a+b*b {
				return a * b * c
			}
		}
	}

	panic("unreachable statement") // This case should never happen.
}
