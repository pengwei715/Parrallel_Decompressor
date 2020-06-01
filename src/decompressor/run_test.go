package main

import (
	"runtime"
	"testing"
)

func TestSeq(t *testing.T) {
	jobs := [3]string{"../../test_data/", "../../test_data/", "../../test_data/"}
	for _, job := range jobs {
		seqworker(job)
	}
}

func TestPar(t *testing.T) {
	runtime.GOMAXPROCS(4)
	jobs := [3]string{"../../test_data/", "../../test_data/", "../../test_data/"}
	for _, job := range jobs {
		seqworker(job)
	}
}

func benchmarkSeq(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		seqworker("../../test_data/")
	}
}

func BenchmarkSeq1(b *testing.B)  { benchmarkSeq(1, b) }
func BenchmarkSeq2(b *testing.B)  { benchmarkSeq(2, b) }
func BenchmarkSeq3(b *testing.B)  { benchmarkSeq(3, b) }
func BenchmarkSeq10(b *testing.B) { benchmarkSeq(10, b) }
func BenchmarkSeq20(b *testing.B) { benchmarkSeq(20, b) }
func BenchmarkSeq40(b *testing.B) { benchmarkSeq(40, b) }

func benchmarkPar(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		parworker("../../test_data/")
	}
}

func BenchmarkPar1(b *testing.B)  { benchmarkPar(1, b) }
func BenchmarkPar2(b *testing.B)  { benchmarkPar(2, b) }
func BenchmarkPar3(b *testing.B)  { benchmarkPar(3, b) }
func BenchmarkPar10(b *testing.B) { benchmarkPar(10, b) }
func BenchmarkPar20(b *testing.B) { benchmarkPar(20, b) }
func BenchmarkPar40(b *testing.B) { benchmarkPar(40, b) }
