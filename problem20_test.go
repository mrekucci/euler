// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import (
	"math/big"
	"testing"
)

func TestProblem20(t *testing.T) {
	want := 648
	got := Problem20()
	if got != want {
		t.Errorf("Problem20() = %d; want %d", got, want)
	}
}

var factorialTests = []struct {
	in  int64
	out int64
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
		got := factorial(big.NewInt(tt.in))
		if got.Cmp(big.NewInt(tt.out)) != 0 {
			t.Errorf("factorial(%d) = %v; want %d", tt.in, got, tt.out)
		}
	}
}

func BenchmarkProblem20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem20()
	}
}
