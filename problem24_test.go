// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem24(t *testing.T) {
	const want = 2783915460
	got, err := Problem24()
	if got != want || err != nil {
		t.Errorf("Problem24() = %d, %v; want %d, <nil>", got, err, want)
	}
}

var factorialTests = []struct {
	in   int
	want int
}{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{5, 120},
}

func TestFactorial(t *testing.T) {
	for _, tt := range factorialTests {
		got := factorial(tt.in)
		if got != tt.want {
			t.Errorf("factorial(%d) = %v; want %d", tt.in, got, tt.want)
		}
	}
}

func BenchmarkProblem24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem24()
	}
}
