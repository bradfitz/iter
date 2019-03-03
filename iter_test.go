// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iter_test

import (
	"fmt"
	"testing"

	"github.com/bradfitz/iter"
)

func ExampleN() {
	for i := range iter.N(4) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
}

func TestAllocs(t *testing.T) {
	var sink []struct{}
	allocs := testing.AllocsPerRun(1000, func() {
		sink = iter.N(1e9)
	})
	_ = sink
	if allocs > 0.1 {
		t.Errorf("allocs = %v", allocs)
	}
}
