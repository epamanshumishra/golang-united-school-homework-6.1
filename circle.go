package golang_united_school_homework

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

// CalcPerimeter returns calculation result of perimeter
func (c Circle) CalcPerimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// CalcArea returns calculation result of area
func (c Circle) CalcArea() float64 {
	return 3.14 * c.Radius * c.Radius
}
