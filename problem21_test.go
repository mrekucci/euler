// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem21(t *testing.T) {
	const want = 31626
	got := Problem21()
	if got != want {
		t.Errorf("Problem21() = %d; want %d", got, want)
	}
}

var isAmicableTests = []struct {
	in  int
	out bool
}{
	{1, false},
	{220, true},
	{284, true},
	{300, false},
	{1328470, true},
	{1483850, true},
	{2000000, false},
	{8619765, true},
	{9627915, true},
}

func TestIsAmicable(t *testing.T) {
	for _, tt := range isAmicableTests {
		if isAmicable(tt.in) != tt.out {
			t.Errorf("isAmicable(%d) = %t; want %t", tt.in, !tt.out, tt.out)
		}
	}
}

func BenchmarkProblem21(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem21()
	}
}
