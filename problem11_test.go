// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem11(t *testing.T) {
	const want = 70600674
	got := Problem11()
	if got != want {
		t.Errorf("Problem11() = %d; want %d", got, want)
	}
}

func BenchmarkProblem11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem11()
	}
}
