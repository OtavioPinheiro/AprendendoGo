package main

import "testing"

type test struct {
	dados []int
	resp  int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		test{dados: []int{1, 2, 3}, resp: 6},
		test{dados: []int{-1, 1, 0}, resp: 0},
		test{dados: []int{-10, -1}, resp: -11},
		test{dados: []int{0, 0, 0}, resp: 0},
	}
	for _, v := range tests {
		z := Soma(v.dados...)
		if z != v.resp {
			t.Error("Esperava:", v.resp, "Recebi:", z)
		}
	}
}
