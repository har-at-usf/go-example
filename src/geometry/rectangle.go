package geometry

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perim() float64 {
	return 2*r.Width + 2*r.Height
}
