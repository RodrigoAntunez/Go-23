package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)
	if total != 10 {
		t.Errorf("Suma de 5 + 5 debe ser 10, pero fue %d", total)
	}
}
