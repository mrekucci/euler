// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

// Problem19 is solution for finding how many Sundays fell on the first
// of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000).
func Problem19() int {
	// First solution base on time package:
	//
	//	cnt := 0
	//	for y := 1901; y <= 2000; y++ {
	//		for m := 1; m <= 12; m++ {
	//			if time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC).Weekday() == time.Sunday {
	//				cnt++
	//			}
	//		}
	//	}
	//	return cnt
	//

	// Second solution without using time package, longer but faster.
	months := []int{
		31, // Jan
		28, // Feb
		31, // Mar
		30, // Apr
		31, // May
		30, // Jun
		31, // Jul
		31, // Aug
		30, // Sep
		31, // Oct
		30, // Nov
		31, // Dec
	}

	cnt := 0
	day := 2 // Represents the day 0-6 => Sun to Sat. We start with '2' because it was Tuesday on 1/1/1901.
	for y := 1901; y <= 2000; y++ {
		ly := isLeapYear(y)
		for m, d := range months {
			if m == 1 && ly { // Check if Feb. should have 29 days.
				d++
			}

			day += (d % 28) // Add remaining days of month, this shift us to the end of the month.
			// The actual day can be computed as day%7, 'cause %7 generates day from 0 to 6
			// day+1 shift us at the first day of the next month.
			if (day+1)%7 == 0 {
				cnt++
				// We can reset the day (day %= 7) to 0-6 after statement cnt++
				// But this isn't necessary, because we care only about the
				// remainder. It will be always the same with or without the reset.
			}
		}
	}
	return cnt
}

// isLeapYear returns true if y is a leap year.
func isLeapYear(y int) bool {
	if y%400 == 0 || y%4 == 0 && !(y%100 == 0) {
		return true
	}
	return false
}
