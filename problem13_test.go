// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem13(t *testing.T) {
	const want = 5537376230
	got, err := Problem13()
	if got != want || err != nil {
		t.Errorf("Problem13() = %d, %v; want %d, <nil>", got, err, want)
	}
}

func BenchmarkProblem13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem13()
	}
}
