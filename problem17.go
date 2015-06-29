// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

// Problem17 is solution for finding how many letters would
// be used in all the numbers from 1 to 1000 (one thousand)
// inclusive written out in words without spaces or hyphens.
func Problem17() int {
	cnt := 0

	units := []int{
		len("one"),
		len("two"),
		len("three"),
		len("four"),
		len("five"),
		len("six"),
		len("seven"),
		len("eight"),
		len("nine"),
	}

	teens := []int{
		len("ten"),
		len("eleven"),
		len("twelve"),
		len("thirteen"),
		len("fourteen"),
		len("fifteen"),
		len("sixteen"),
		len("seventeen"),
		len("eighteen"),
		len("nineteen"),
	}

	tens := []int{
		len("twenty"),
		len("thirty"),
		len("forty"),
		len("fifty"),
		len("sixty"),
		len("seventy"),
		len("eighty"),
		len("ninety"),
	}

	// Add 1-9.
	for _, u := range units {
		cnt += u
	}

	// Add 10-19.
	for _, t := range teens {
		cnt += t
	}

	// Add 20-99.
	for _, t := range tens {
		cnt += t
		for _, u := range units {
			cnt += t + u
		}
	}

	// Add 100-999.
	nn := cnt // Length of numbers 1-99.
	h := len("hundred")
	for _, u := range units {
		// 99*(u+h+3) replace cycle for doing this 99 times on first 99 numbers.
		// E.g: "one hundred and twenty-two" (without spaces and hyphens).
		cnt += (u + h) + (99*(u+h+3) + nn)
	}

	// Add 1000.
	cnt += len("onethousand")

	return cnt
}
