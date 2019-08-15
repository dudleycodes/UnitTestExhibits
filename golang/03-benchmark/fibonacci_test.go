package fibonacci

import (
	"fmt"
	"testing"
)

func Test_Fib(t *testing.T) {
	tests := map[uint]uint{0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 10: 55, 42: 267914296}

	for input, expected := range tests {
		t.Run(fmt.Sprintf("%d", input), func(t *testing.T) {
			actual := fib(input)
			if actual != expected {
				t.Errorf("With an input of %d expected an output of %d but got %d.", input, expected, actual)
			}
		})
	}
}

func Benchmark_Fib(b *testing.B) {
	b.Run("Calculate 20", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			fib(20) // run the Fib function b.N times
		}
	})
}
