// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import "testing"

func TestAssertEquals(t *testing.T) {
	v := 1
	AssertEquals(t, v, v)
	if t.Failed() {
		t.Errorf("AssertEquals(t, %d, %d) fails; want %d equals %d", v, v, v, v)
	}
}
