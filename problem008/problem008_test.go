// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem008

import "testing"

func TestSolution(t *testing.T) {
	want := 23514624000
	got, err := Solution()
	if got != want || err != nil {
		t.Errorf("Solution = %d, %v; want %d, %v", got, err, want, nil)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution()
	}
}
