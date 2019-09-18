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
		perimetr() float64
}

func totalArea(shapes ...Shape) float64 {
		var area float64
		for _, s := range shapes {
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

func (r *Rectangle) perimetr() float64 {
		l := distance(r.x1, r.y1, r.x1, r.y2)
		w := distance(r.x1, r.y1, r.x2, r.y1)
		return 2*l + 2*w
}

func (c *Circle) perimetr() float64 {
		return math.Pi * c.r*2
}

func main() {
		r := Rectangle{0, 0, 10, 10}
		fmt.Println(r.area())

		c := Circle{0, 0, 5}
		fmt.Println(c.area())

		// use Shape interface
		fmt.Println(totalArea(&c, &r))

		// Perimetr
		fmt.Println("P of rectangle:", r.perimetr())
		fmt.Println("P of circle:", c.perimetr())
}