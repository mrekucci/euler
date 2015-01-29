// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem9(t *testing.T) {
	want := 31875000
	got := Problem9()
	if got != want {
		t.Errorf("Problem9() = %d; want %d", got, want)
	}
}

func BenchmarkProblem9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem9()
	}
}
