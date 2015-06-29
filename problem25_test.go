// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem25(t *testing.T) {
	const want = 4782
	got := Problem25()
	if got != want {
		t.Errorf("Problem25() = %d; want %d", got, want)
	}
}

func BenchmarkProblem25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem25()
	}
}
