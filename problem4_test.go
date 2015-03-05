// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem4(t *testing.T) {
	const want = 906609
	got := Problem4()
	if got != want {
		t.Errorf("Problem4() = %d; want %d", got, want)
	}
}

var reverseTests = []struct {
	in   int
	want int
}{
	{1111, 1111},
	{1010, 101},
	{1234, 4321},
}

func TestReverse(t *testing.T) {
	for _, tt := range reverseTests {
		got := reverse(tt.in)
		if got != tt.want {
			t.Errorf("reverse(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

var isPalindromeTests = []struct {
	in   int
	want bool
}{
	{101, true},
	{1010, false},
	{12321, true},
	{123456, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range isPalindromeTests {
		got := isPalindrome(tt.in)
		if got != tt.want {
			t.Errorf("isPalindrome(%d) = %t; want %t", tt.in, got, tt.want)
		}
	}
}

func BenchmarkProblem4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem4()
	}
}
