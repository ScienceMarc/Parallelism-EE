package main

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkPrimes(b *testing.B) {
	functions := []struct {
		name string
		fun  (func(x uint64) []uint64)
	}{{"Single Threaded", findPrimes}, {"Multi Threaded", findPrimesParallel}}

	for _, v := range [...]uint64{1, 2, 3, 4, 5, 6, 7, 8} {
		maxVal := uint64(math.Pow(10, float64(v)))
		for _, f := range functions {
			b.Run(fmt.Sprintf("%s%d", f.name, maxVal), func(b *testing.B) {
				for bench := 0; bench < b.N; bench++ {
					f.fun(maxVal)
				}
			})
		}
	}

}
