// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem17(t *testing.T) {
	const want = 21124
	got := Problem17()
	if got != want {
		t.Errorf("Problem17() = %d; want %d", got, want)
	}
}

func BenchmarkProblem17(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem17()
	}
}
