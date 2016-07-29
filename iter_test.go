package iter_test

import (
	"fmt"
	"testing"

	"github.com/voutasaurus/iter"
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
	var x []struct{}
	allocs := testing.AllocsPerRun(500, func() {
		x = iter.N(1e9)
	})
	if allocs > 1.1 {
		t.Errorf("allocs = %v", allocs)
	}
}

func TestNest(t *testing.T) {
	l := len(iter.N(&struct{ N []string }{N: make([]string, 42)}))
	if l != 42 {
		t.Errorf("expected: %d, got: %d", 42, l)
	}
}

func TestNil(t *testing.T) {
	z := len(iter.N(nil))
	if z != 0 {
		t.Errorf("iter.N(nil): expected: 0, got: %d", z)
	}
}

func TestNilPtr(t *testing.T) {
	var x *int
	z := len(iter.N(x))
	if z != 0 {
		t.Errorf("iter.N(nil): expected: 0, got: %d", z)
	}
}

func TestNilInterface(t *testing.T) {
	var x iter.Inter
	z := len(iter.N(x))
	if z != 0 {
		t.Errorf("iter.N(nil): expected: 0, got: %d", z)
	}
}
