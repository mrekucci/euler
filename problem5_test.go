// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem5(t *testing.T) {
	want := 232792560
	got := Problem5()
	if got != want {
		t.Errorf("Problem5() = %d; want %d", got, want)
	}
}

func BenchmarkProblem5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem5()
	}
}
