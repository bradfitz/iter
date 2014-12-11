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
	var x []struct{}
	allocs := testing.AllocsPerRun(500, func() {
		x = iter.N(1e9)
	})
	if allocs > 0.1 {
		t.Errorf("allocs = %v", allocs)
	}
}

// To benchmark, run: go test -bench=.

// Global result variable to prevent compiler from optimizing calls away. It will
// likely overflow but this doesn't matter since the actual value is never used
var result int

func benchIter(n int, b *testing.B) {
	r := 0
	for x := 0; x < b.N; x++ {
		for i := range iter.N(n) {
			r += i
		}
	}
	result = r
}

func benchFor(n int, b *testing.B) {
	r := 0
	for x := 0; x < b.N; x++ {
		for i := 0; i < n; i++ {
			r += i
		}
	}
	result = r
}

func BenchmarkIter1e0(b *testing.B) { benchIter(1e0, b) }
func BenchmarkIter1e1(b *testing.B) { benchIter(1e1, b) }
func BenchmarkIter1e2(b *testing.B) { benchIter(1e2, b) }
func BenchmarkIter1e3(b *testing.B) { benchIter(1e3, b) }
func BenchmarkIter1e4(b *testing.B) { benchIter(1e4, b) }
func BenchmarkIter1e5(b *testing.B) { benchIter(1e5, b) }
func BenchmarkIter1e6(b *testing.B) { benchIter(1e6, b) }

// Function names padded to "01" for better aligned output.
func BenchmarkFor01e0(b *testing.B) { benchFor(1e0, b) }
func BenchmarkFor01e1(b *testing.B) { benchFor(1e1, b) }
func BenchmarkFor01e2(b *testing.B) { benchFor(1e2, b) }
func BenchmarkFor01e3(b *testing.B) { benchFor(1e3, b) }
func BenchmarkFor01e4(b *testing.B) { benchFor(1e4, b) }
func BenchmarkFor01e5(b *testing.B) { benchFor(1e5, b) }
func BenchmarkFor01e6(b *testing.B) { benchFor(1e6, b) }
