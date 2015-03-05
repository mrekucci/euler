// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

func TestProblem19(t *testing.T) {
	const want = 171
	got := Problem19()
	if got != want {
		t.Errorf("Problem19() = %d; want %d", got, want)
	}
}

var isLeapYearTests = []struct {
	in   int
	want bool
}{
	{1904, true},
	{1910, false},
	{1920, true},
	{1930, false},
	{1932, true},
	{1950, false},
	{1955, false},
	{1960, true},
	{1970, false},
	{1980, true},
	{1990, false},
	{2000, true},
}

func TestIsLapYear(t *testing.T) {
	for _, tt := range isLeapYearTests {
		got := isLeapYear(tt.in)
		if got != tt.want {
			t.Errorf("isLeapYear(%d) = %t; want %t", tt.in, got, tt.want)
		}
	}
}

func BenchmarkProblem19(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem19()
	}
}
