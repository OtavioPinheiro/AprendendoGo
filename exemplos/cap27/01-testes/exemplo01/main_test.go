package main

import "testing"

func TestAdicao(t *testing.T) {
	teste := adicao(1, 1)
	esperado := 2
	if teste != esperado {
		t.Errorf("Esperava: %v, mas recebi %v", esperado, teste)
	}
}

func TestFailAdicao(t *testing.T) {
	teste := adicao(1, 1, 3)
	esperado := 2
	if teste != esperado {
		t.Errorf("Esperava: %v, mas recebi %v", esperado, teste)
	}
}
