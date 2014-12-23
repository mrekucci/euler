// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem001

import "fmt"

// Solution for finding the sum of all the multiples of 3 or 5 below 1000.
func Solution() int {
	var s3, s5, s15 int
	var err error

	if s3, err = sn(3, 999); err != nil {
		panic(err)
	}
	if s5, err = sn(5, 999); err != nil {
		panic(err)
	}
	if s15, err = sn(15, 999); err != nil {
		panic(err)
	}
	return s3 + s5 - s15
}

// sn returns the sum of arithmetic sequence
// up to the number an. The sequence has difference d.
// An error is returned when d >= an.
func sn(d, an int) (int, error) {
	if d >= an {
		return 0, fmt.Errorf("illegal argument %d >= %d; want %d < %d", d, an, d, an)
	}

	n := an / d
	return int(float64(n) / 2 * float64((2*d)+(n-1)*d)), nil
}
