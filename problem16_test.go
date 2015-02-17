// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem16(t *testing.T) {
	const want = 1366
	got := Problem16()
	if got != want {
		t.Errorf("Problem16() = %d; want %d", got, want)
	}
}

func BenchmarkProblem16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem16()
	}
}
