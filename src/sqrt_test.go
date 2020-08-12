package main

import "testing"

func TestSqrt(t *testing.T) {
	sqrt := Sqrt(0.0001, 10000000)
	if sqrt != "99999.999986" {
		t.Errorf("Resposta incorreta. Chegou: %s, deveria: %s.", sqrt, "99999.999986")
	}
}
