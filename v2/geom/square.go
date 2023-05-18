package geom

type Square struct {
	side float64
}

func (s Square) Perimeter() float64 {
	return s.side * 4.0
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func SumAreas(a float64, b float64) float64 {
	return a + b
}
