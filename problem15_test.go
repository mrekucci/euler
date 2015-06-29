// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem15(t *testing.T) {
	const want = 137846528820
	got := Problem15()
	if got != want {
		t.Errorf("Problem15() = %d; want %d", got, want)
	}
}

func BenchmarkProblem15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem15()
	}
}
