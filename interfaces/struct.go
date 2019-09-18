package main

import "fmt"

type Circle struct {
		x, y, r float64
}

type Rectangle struct {
		x1 float64
		x2 float64
		y1 float64
		y2 float64
}

func main() {
		// create Structs
		var c Circle
		r1 := Rectangle{x1: 1, x2: 2, y1: 3, y2: 5}
		r2 := Rectangle{1, 2, 3, 5}

		// show Structs
		fmt.Println(c)
		fmt.Println(r1)
		fmt.Println(r2)

		// show Struct fields
		fmt.Println(r1.x1, r2.x2)
}

