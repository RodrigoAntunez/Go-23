package main

type Circle interface {
	area() float64
	perimetro() float64
}

type Square interface {
	area() float64
	perimetro() float64
}

type square struct {
	sideLength float64
}

func (s square) Area() float64 {
	return s.sideLength * s.sideLength
}

func (s square) Perimetro() float64 {
	return s.sideLength * 4
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius

}

func (c circle) Perimetro() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	// mySquare := square{sideLength: 10}
	// area := mySquare.Area()
	// perimetro := mySquare.Perimetro()

	// myCircle := circle{radius: 10}
	// area := myCircle.Area()
	// perimetro := myCircle.Perimetro()
}
