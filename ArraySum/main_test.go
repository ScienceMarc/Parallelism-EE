package main

import (
	"testing"
)

//TODO: Write better benchmarks to figure out WTF is going on

func benchmarkAddMulticore(size int, b *testing.B) {
	longArray := make([]int, 1<<size)
	for i := 0; i < 1<<size; i++ {
		longArray[i] = 1
	}
	//fmt.Println("allocated array")
	for bench := 0; bench < b.N; bench++ {
		multiSum(longArray)
	}
}

/*
func BenchmarkAddMulticore5(b *testing.B)  { benchmarkAddMulticore(5, b) }
func BenchmarkAddMulticore10(b *testing.B) { benchmarkAddMulticore(10, b) }
func BenchmarkAddMulticore15(b *testing.B) { benchmarkAddMulticore(15, b) }*/
func BenchmarkAddMulticore20(b *testing.B) { benchmarkAddMulticore(20, b) } /*
func BenchmarkAddMulticore25(b *testing.B) { benchmarkAddMulticore(25, b) }

//func BenchmarkAddMulticore30(b *testing.B) { benchmarkAddMulticore(30, b) }

func benchmarkAddSinglethread(size int, b *testing.B) {
	longArray := make([]int, 1<<size)
	for i := 0; i < 1<<size; i++ {
		longArray[i] = 1
	}
	for bench := 0; bench < b.N; bench++ {
		sum(longArray)
	}
}

func BenchmarkAddSinglethread5(b *testing.B)  { benchmarkAddSinglethread(5, b) }
func BenchmarkAddSinglethread10(b *testing.B) { benchmarkAddSinglethread(10, b) }
func BenchmarkAddSinglethread15(b *testing.B) { benchmarkAddSinglethread(15, b) }
func BenchmarkAddSinglethread20(b *testing.B) { benchmarkAddSinglethread(20, b) }
func BenchmarkAddSinglethread25(b *testing.B) { benchmarkAddSinglethread(25, b) }

//func BenchmarkAddSinglethread30(b *testing.B) { benchmarkAddSinglethread(30, b) }
*/
