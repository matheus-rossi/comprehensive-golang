package main

import "fmt"

type rectangle struct {
	height float32
	width  float32
}

func (r rectangle) area() float32 {
	return r.height * r.width
}

func (r rectangle) perimeter() float32 {
	return 2*r.height + 2*r.width
}

func main() {
	rectangle := rectangle{height: 10, width: 20}

	x_area := rectangle.area()
	x_perimeter := rectangle.perimeter()

	fmt.Printf("Area = %v \n", x_area)
	fmt.Printf("Perimetro = %v \n", x_perimeter)
}
