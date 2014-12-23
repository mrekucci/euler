// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "testing"

// AssertEquals fails by Errorf if res != want.
// It helps reduce boilerplate code in tests.
func AssertEquals(t *testing.T, res int, want int) {
	if res != want {
		t.Errorf("Result = %d; want %d", res, want)
	}
}
