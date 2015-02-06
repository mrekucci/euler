// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem21(t *testing.T) {
	want := 31626
	got := Problem21()
	if got != want {
		t.Errorf("Problem21() = %d; want %d", got, want)
	}
}

var isAmicableTests = []struct {
	in  int
	out bool
}{
	{1, false},
	{220, true},
	{284, true},
	{300, false},
	{1328470, true},
	{1483850, true},
	{2000000, false},
	{8619765, true},
	{9627915, true},
}

func TestIsAmicable(t *testing.T) {
	for _, tt := range isAmicableTests {
		if isAmicable(tt.in) != tt.out {
			t.Errorf("isAmicable(%d) = %t; want %t", tt.in, !tt.out, tt.out)
		}
	}
}

var propDivSumTests = []struct {
	in  int
	out int
}{
	{0, 0},
	{1, 0},
	{4, 3},
	{16, 15},
	{220, 284},
	{284, 220},
	{8262136, 8369864},
	{18655744, 19154336},
}

func TestPropDivSum(t *testing.T) {
	for _, tt := range propDivSumTests {
		got := propDivSum(tt.in)
		if got != tt.out {
			t.Errorf("propDivSum(%d) = %d; want %d", tt.in, got, tt.out)
		}
	}
}

func BenchmarkProblem21(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem21()
	}
}
