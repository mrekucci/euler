// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "math/big"

var one = big.NewInt(1)

// Problem20 is solution for finding the sum of the digits in the number 100!
func Problem20() int {
	// First solution without using custom factorial function:
	//
	//	sum := 0
	//	for _, d := range new(big.Int).MulRange(1, 100).String() {
	//		sum += int(d - '0')
	//	}
	//	return sum
	//

	// Second solution using custom factorial function is a bit faster:
	sum := 0
	for _, d := range factorial(big.NewInt(100)).String() {
		sum += int(d - '0')
	}
	return sum
}

// factorial computes factorial of number n.
// This version is memory less expensive and
// faster like its next recursive version:
//
//	func factorial(n *big.Int) *big.Int {
// 		if n.Cmp(one) == 1 {
// 			// new(big.Int) 'cause we will modify 'n' which is a *big.Int.
// 			return n.Mul(n, factorial(new(big.Int).Sub(n, one)))
// 		}
//		return one
//	}
//
func factorial(n *big.Int) *big.Int {
	if n.Cmp(one) == 1 {
		// new(big.Int) 'cause we will modify 'n' which is a *big.Int.
		cn := new(big.Int).Set(n)
		for i := big.NewInt(1); i.Cmp(cn) != 0; i.Add(i, one) {
			n.Mul(n, i) // Same as n*=i
		}
		return n
	}
	return one
}
