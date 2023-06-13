package src

import "fmt"

type figure2d interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	heigth float64
}

func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.base * r.heigth
}

func area(f figure2d) string {
	return fmt.Sprintf("Area result: %f", f.area())
}

func TestInterfaces() {
	s := square{base: 4}
	r := rectangle{base: 2, heigth: 4}

	fmt.Println(area(s))
	fmt.Println(area(r))
}
