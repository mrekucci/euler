// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem8(t *testing.T) {
	const want = 23514624000
	got, err := Problem8()
	if got != want || err != nil {
		t.Errorf("Problem8() = %d, %v; want %d, <nil>", got, err, want)
	}
}

func BenchmarkProblem8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem8()
	}
}
