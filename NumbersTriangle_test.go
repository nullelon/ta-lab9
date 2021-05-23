package main

import "testing"

type function func(nt *NumbersTriangle) *Solution

func doBench(f function, height int, b *testing.B) {
	n := height * (height + 1) / 2
	arr := make([]int, n)
	nt := newNumbersTriangle(arr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(nt)
	}
}

func BenchmarkCalculateTriangleMaxSumWithDP_5(b *testing.B) {
	doBench(calculateTriangleMaxSumWrap, 5, b)
}
func BenchmarkCalculateTriangleMaxSumWithDP_10(b *testing.B) {
	doBench(calculateTriangleMaxSumWrap, 10, b)
}
func BenchmarkCalculateTriangleMaxSumWithDP_15(b *testing.B) {
	doBench(calculateTriangleMaxSumWrap, 15, b)
}
func BenchmarkCalculateTriangleMaxSumWithDP_20(b *testing.B) {
	doBench(calculateTriangleMaxSumWrap, 20, b)
}
func BenchmarkCalculateTriangleMinSumWithDP_5(b *testing.B) {
	doBench(calculateTriangleMinSumWrap, 5, b)
}
func BenchmarkCalculateTriangleMinSumWithDP_10(b *testing.B) {
	doBench(calculateTriangleMinSumWrap, 10, b)
}
func BenchmarkCalculateTriangleMinSumWithDP_15(b *testing.B) {
	doBench(calculateTriangleMinSumWrap, 15, b)
}
func BenchmarkCalculateTriangleMinSumWithDP_20(b *testing.B) {
	doBench(calculateTriangleMinSumWrap, 20, b)
}

//
func BenchmarkCalculateTriangleMaxSum_NoDP_5(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 5, b)
}
func BenchmarkCalculateTriangleMaxSum_NoDP_10(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 10, b)
}
func BenchmarkCalculateTriangleMaxSum_NoDP_15(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 15, b)
}
func BenchmarkCalculateTriangleMaxSum_NoDP_20(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 20, b)
}
func BenchmarkCalculateTriangleMinSum_NoDP_5(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 5, b)
}
func BenchmarkCalculateTriangleMinSum_NoDP_10(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 10, b)
}
func BenchmarkCalculateTriangleMinSum_NoDP_15(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 15, b)
}
func BenchmarkCalculateTriangleMinSum_NoDP_20(b *testing.B) {
	doBench(calculateTriangleMinSum_noDPWrap, 20, b)
}
