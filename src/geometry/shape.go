package geometry

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perim() float64
}

func Measure(g Shape) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}
