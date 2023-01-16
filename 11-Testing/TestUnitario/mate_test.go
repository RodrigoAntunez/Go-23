package testunitario

import "testing"

//TEST UNITARIO
/*func TestSuma(t *testing.T) {
	total := Suma(5, 5)
	if total != 10 {
		t.Errorf("Suma de 5 + 5 debe ser 10, pero fue %d", total)
	}
} */

//TEST DE TABLAS (Convencional)

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}
	for _, item := range tabla {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("Suma de %d + %d debe ser %d, pero fue %d", item.a, item.b, item.c, total)
		}
	}
}
