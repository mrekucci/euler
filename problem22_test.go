// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem22(t *testing.T) {
	const want = 871198282
	got, err := Problem22()
	if got != want || err != nil {
		t.Errorf("Problem22() = %d, %v; want %d, <nil>", got, err, want)
	}
}

func BenchmarkProblem22(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem22()
	}
}
