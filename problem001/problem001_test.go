// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem001

import (
	"testing"

	"github.com/mrekucci/euler/util"
)

func TestSolution(t *testing.T) {
	util.AssertEquals(t, Solution(), 233168)
}

func TestSn(t *testing.T) {
	d, an := 1, 5
	r, err := sn(d, an)
	if err != nil {
		t.Fatalf("sn(%d, %d) returns error: %s", d, an, err)
	}
	util.AssertEquals(t, r, 15)
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution()
	}
}
