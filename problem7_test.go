// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem7(t *testing.T) {
	want := 104743
	got := Problem7()
	if got != want {
		t.Errorf("Problem7() = %d; want %d", got, want)
	}
}

func BenchmarkProblem7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem7()
	}
}
