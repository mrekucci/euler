// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "testing"

func TestAssertEquals(t *testing.T) {
	v := 1
	AssertEquals(t, v, v)
	if t.Failed() {
		t.Errorf("AssertEquals(t, %d, %d) fails; want %d equals %d", v, v, v, v)
	}
}

var isPrimeTests = []struct {
	in  int
	out bool
}{
	{-4, false},
	{-3, false},
	{-2, false},
	{-1, false},
	{0, false},
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{6, false},
	{7, true},
	{11, true},
	{12, false},
	{15, false},
	{17, true},
	{23, true},
	{24, false},
	{25, false},
	{29, true},
	{31, true},
	{100, false},
	{105, false},
}

func TestIsPrime(t *testing.T) {
	for _, tt := range isPrimeTests {
		if IsPrime(tt.in) != tt.out {
			t.Errorf("IsPrime(%d) = %t; want %t", tt.in, !tt.out, tt.out)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(i)
	}
}
