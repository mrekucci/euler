// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "testing"

const (
	noError  = true
	hasError = false
)

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

var mulTests = []struct {
	a   int
	b   int
	out int
	ok  bool
}{
	{MinInt, 0, 0, true},
	{0, MinInt, 0, true},
	{MinInt, 1, MinInt, true},
	{1, MinInt, MinInt, true},
	{MaxInt, -1, -MaxInt, true},
	{-1, MaxInt, -MaxInt, true},
	{-MaxInt, -1, MaxInt, true},
	{-1, -MaxInt, MaxInt, true},
	{MinInt, 2, 0, false},
	{MinInt, -2, 0, false},
	{MinInt, -1, 0, false},
	{2, MinInt, 0, false},
	{-2, MinInt, 0, false},
	{-1, MinInt, 0, false},
}

func TestMul(t *testing.T) {
	for _, tt := range mulTests {
		got, ok := Mul(tt.a, tt.b)
		if got != tt.out || ok != tt.ok {
			t.Errorf("Mul(%d, %d) = %d, %t; want %d, %t", tt.a, tt.b, got, ok, tt.out, tt.ok)
		}
	}
}

var mulIntsTests = []struct {
	in  []int
	out int
	ok  bool
}{
	{[]int{-2, 3, 4}, -24, noError},
	{[]int{-3, -4, 5}, 60, noError},
	{[]int{-4, -5, -6}, -120, noError},
	{[]int{0, 0, 0}, 0, noError},
	{[]int{2, 3, 4}, 24, noError},
	{[]int{3, 4, 5}, 60, noError},
	{[]int{4, 5, 6}, 120, noError},
	{[]int{MinInt, 2}, 0, hasError},
	{[]int{MinInt, -2}, 0, hasError},
	{[]int{MinInt, -1}, 0, hasError},
	{[]int{2, MinInt}, 0, hasError},
	{[]int{-2, MinInt}, 0, hasError},
	{[]int{-1, MinInt}, 0, hasError},
}

func TestMulInts(t *testing.T) {
	for _, tt := range mulIntsTests {
		got, err := MulInts(tt.in)
		switch {
		case got != tt.out && tt.ok:
			t.Errorf("MulInts(%d) = %d, %v; want %d, <nil>", tt.in, got, err, tt.out)
			continue
		case err == nil && !tt.ok:
			t.Errorf("MulInts(%d) = %d, %v; expected error, got none", tt.in, got, err)
			continue
		case err != nil && tt.ok:
			t.Errorf("MulInts(%d) = %d, %v; unexpected error: %v", tt.in, got, err, err)
			continue
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(i)
	}
}

func BenchmarkMulInts(b *testing.B) {
	ints := make([]int, 1e3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MulInts(ints)
	}
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(i, i)
	}
}
