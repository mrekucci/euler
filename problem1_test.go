// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem1(t *testing.T) {
	want := 233168
	got := Problem1()
	if got != want {
		t.Errorf("Problem1() = %d; want %d", got, want)
	}
}

var snTests = []struct {
	d, an int
	out   int
}{
	{1, 5, 15},
	{2, 10, 30},
	{3, 15, 45},
	{4, 1000, 125500},
	{5, 1000, 100500},
	{10, 1000, 50500},
}

func TestSn(t *testing.T) {
	for _, tt := range snTests {
		got := sn(tt.d, tt.an)
		if got != tt.out {
			t.Errorf("sn(%d, %d) = %d; want %d", tt.d, tt.an, got, tt.out)
		}
	}
}

func BenchmarkProblem1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem1()
	}
}
