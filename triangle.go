package golang_united_school_homework

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

// CalcPerimeter returns calculation result of perimeter
func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

// CalcArea returns calculation result of area
func (t Triangle) CalcArea() float64 {
	return 0.433 * t.Side * t.Side
}
