package geom

type Square struct {
	Side float64
}

func (s Square) Perimeter() float64 {
	return s.Side * 4.0
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func SumAreas(a float64, b float64) float64 {
	return a + b
}
