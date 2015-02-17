// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem14(t *testing.T) {
	const want = 837799
	got := Problem14()
	if got != want {
		t.Errorf("Problem14() = %d; want %d", got, want)
	}
}

var collatzTests = []struct {
	in  int
	out int
}{
	{13, 40},
	{40, 20},
	{20, 10},
	{5, 16},
	{16, 8},
	{8, 4},
	{4, 2},
	{2, 1},
}

func TestCollatz(t *testing.T) {
	for _, tt := range collatzTests {
		got := collatz(tt.in)
		if got != tt.out {
			t.Errorf("collatz(%d) = %d; want %d", tt.in, got, tt.out)
		}
	}
}

func BenchmarkProblem14(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem14()
	}
}
