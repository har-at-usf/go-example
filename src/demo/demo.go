package main

import "geometry"

func main() {

	r := geometry.Rectangle{Width: 3, Height: 4}
	c := geometry.Circle{Radius: 5}

	geometry.Measure(r)
	geometry.Measure(c)
}
