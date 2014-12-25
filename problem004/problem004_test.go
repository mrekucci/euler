// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package problem004

import (
	"testing"

	"github.com/mrekucci/euler/util"
)

func TestSolution(t *testing.T) {
	util.AssertEquals(t, Solution(), 906609)
}

func TestReverse(t *testing.T) {
	util.AssertEquals(t, reverse(1234), 4321)
}

func TestIsPalindrome(t *testing.T) {
	p := 101
	if !isPalindrome(p) {
		t.Errorf("isPalindrome(%d) = false; want true", p)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution()
	}
}
