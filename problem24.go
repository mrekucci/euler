// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "strconv"

// Problem24 is solution for finding the millionth lexicographic
// permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9.
func Problem24() (int, error) {
	digits := []byte("0123456789")

	// There is 10! permutations for list 0-9, which is more than 1e6.
	// For lexicographical order if we put a digit 0 first, there's 9!
	// permutations for the rest digits in the list 1-9. For this permutation
	// (9!) we have to find a index a, where 9!*a < 1e6. A digit on index a
	// (list is indexed from 1) from the list 1-9 is the first digit of millionth
	// lexicographic permutation. Next we replace 1e6 by 1e6-9!*a and continue
	// to search for number b (index for second digit), where 8!*b < (1e6-9!*a)...
	// and so on. This can be written as:
	//
	// 1e6 = 9!*digits[a] + 8!*digits[b] + 7!*digits[c] + 6!*digits[d] + 5!*digits[e] +
	//       4!*digits[f] + 3!*digits[g] + 2!*digits[h] + 1!*digits[i] + 0!*digits[j].
	//
	var p []byte
	pn := 1000000 // 1e6
	for n := len(digits) - 1; n >= 0; n-- {
		f := factorial(n)
		i := 1
		for i*f < pn {
			i++
		}
		i-- // -1 'cause i*f must be less then 'pn'.
		pn -= (i * f)
		p = append(p, digits[i])
		digits = append(digits[:i], digits[i+1:]...) // Cut the digit from slice.
	}
	return strconv.Atoi(string(p))
}

// factorial computes factorial of number n.
// For n < 0 the result is not specified.
// The overflow isn't check!
func factorial(n int) int {
	f := 1
	if n > 1 {
		for n > 1 {
			f *= n
			n--
		}
	}
	return f
}
