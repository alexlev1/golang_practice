package main

import ("fmt"; "math")

// struct
type Circle struct {
		x, y, r float64
}

// struct
type Rectangle struct {
		x1, y1, x2, y2 float64
}

// interface
type Shape interface {
		area() float64
}

func totalArea(shapes ...Shape) float64 {
		var area float64
		for _, s := range shapes {
				area += s.area()
		}
		return area
}

// interface with fields
type MultiShape struct {
		shapes []Shape
}

func (m *MultiShape) area() float64 {
		var area float64
		for _, s := range m.shapes {
				area += s.area()
		}
		return area
}

func distance(x1, y1, x2, y2 float64) float64 {
		a := x2 - x1
		b := y2 - y1
		return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
		l := distance(r.x1, r.y1, r.x1, r.y2)
		w := distance(r.x1, r.y1, r.x2, r.y1)
		return l * w
}

func (c *Circle) area() float64 {
		return math.Pi * c.r*c.r
}

func main() {
		r := Rectangle{0, 0, 10, 10}
		fmt.Println(r.area())

		c := Circle{0, 0, 5}
		fmt.Println(c.area())

		// use Shape interface
		fmt.Println(totalArea(&c, &r))

		// use MultiShape interface
		var mshape MultiShape
		mshape.shapes = []Shape{&r, &c}
		fmt.Println(mshape.area())
}