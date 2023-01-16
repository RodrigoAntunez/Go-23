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

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{5, 2, 5},
		{10, 4, 10},
		{2, 4, 4},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)

		if max != item.c {
			t.Errorf("El mayor entre %d y %d debe ser %d, pero fue %d", item.a, item.b, item.c, max)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		fibo := Fibonacci(item.n)

		if fibo != item.r {
			t.Errorf("El fibonacci de %d debe ser %d, pero fue %d", item.n, item.r, fibo)
		}
	}
}
