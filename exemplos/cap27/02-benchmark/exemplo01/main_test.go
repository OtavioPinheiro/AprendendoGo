package main

import "testing"

type test struct {
	dados []int
	resp  int
}

func BenchmarkSoma(b *testing.B) {
	tests := []test{
		test{dados: []int{1, 2, 3}, resp: 6},
		test{dados: []int{-1, 1, 0}, resp: 0},
		test{dados: []int{-10, -1}, resp: -11},
		test{dados: []int{0, 0, 0}, resp: 0},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			Soma(v.dados...)
		}
	}
}

func BenchmarkMultiplicacao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplicacao(5, 5)
	}
}
