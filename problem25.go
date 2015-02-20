// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "math"

// Problem25 is solution for finding the first term
// in the Fibonacci sequence to contain 1000 digits.
func Problem25() int {
	// The first solution based on brut force:
	//
	//	term := 3 // 3 'cause we set c to a + b so it is 3th term.
	//	a := big.NewInt(1)
	//	b := big.NewInt(1)
	//	c := big.NewInt(2) // c = a + b.
	//	for len(c.String()) < 1000 {
	//		a.Add(b, c)
	//		b.Add(c, a)
	//		c.Add(a, b)
	//		term += 3 // +3 'cause we are moving by tree: a, b, c, a, b, c, ...
	//	}
	//	return term
	//

	// Second solution based on math formula:
	// Fibonacci(n) = (φ^n−(1−φ)^n)/√5.
	// The number of digits of any number is:
	// d = 1+log_b(x), where b is the base (10 in our case).
	// So d = 1 + log10((φ^n − (1−φ)^n) / √5) => 1 + log10((φ^n − (1−φ)^n)) - log10(√5).
	// If n goes to ∞, (1−φ)^n goes to 0 'cause 1−φ<|1|.
	// So this term is negligible for big n (like ours), and we can rewrite the
	// formula for d to: d = 1 + n*log10(φ) - log10(√5).
	// From which n = (d - 1 + log10(√5)) / log10(φ), n must be ceil.
	//
	return int(math.Ceil((1000 - 1 + math.Log10(math.Sqrt(5))) / math.Log10(math.Phi)))
}
