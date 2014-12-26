// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem006

import "math"

// Solution for finding the difference between the sum of the squares
// of the first one hundred natural numbers and the square of the sum.
func Solution() int {
	const n float64 = 100
	sumSqr := (n / 6) * (n + 1) * (2*n + 1)    // Sum of arithmetic sequence: i^2; from i=1 to i=n.
	sqrSum := math.Pow((n/2)*(2*1+(n-1)*1), 2) // Power of Sum of arithmetic sequence: i; from i=1 to i=n.
	return int(sqrSum - sumSqr)
}
