// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem12(t *testing.T) {
	want := 76576500
	got := Problem12()
	if got != want {
		t.Errorf("Problem12() = %d; want %d", got, want)
	}
}

func BenchmarkProblem12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem12()
	}
}
