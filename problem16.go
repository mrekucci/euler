// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "math/big"

// Problem16 is solution for finding the sum of the digits of the number 2^1000.
func Problem16() int {
	num := big.NewInt(1).Exp(big.NewInt(2), big.NewInt(1000), nil).String()
	sum := 0
	for _, d := range num {
		sum += int(d - '0')
	}
	return sum
}
