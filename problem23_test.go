// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem23(t *testing.T) {
	const want = 4179871
	got := Problem23()
	if got != want {
		t.Errorf("Problem23() = %d; want %d", got, want)
	}
}

var isAbundantTests = []struct {
	in  int
	out bool
}{
	{0, false},
	{10, false},
	{12, true},
	{50, false},
	{54, true},
	{81, false},
	{100, true},
	{105, false},
	{120, true},
}

func TestIsAbundant(t *testing.T) {
	for _, tt := range isAbundantTests {
		if isAbundant(tt.in) != tt.out {
			t.Errorf("isAbundant(%d) = %t; want %t", tt.in, !tt.out, tt.out)
		}
	}
}

func BenchmarkProblem23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem23()
	}
}
