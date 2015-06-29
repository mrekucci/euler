// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

import "testing"

func TestProblem6(t *testing.T) {
	const want = 25164150
	got := Problem6()
	if got != want {
		t.Errorf("Problem6() = %d; want %d", got, want)
	}
}

func BenchmarkProblem6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem6()
	}
}
