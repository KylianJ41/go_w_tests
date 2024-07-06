package structsmethodinterfaces

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.141592653589793
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	//return 0.5 * t.Base * t.Height
	return 0
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}
