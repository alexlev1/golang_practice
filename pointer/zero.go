package main

import "fmt"

func zero(x *int) {
		*x = 0
}

func one(y *int) {
		*y = 1
}

func main() {
		x := 5
		zero(&x)
		fmt.Println(x) // x is 0

		y := new(int)
		one(y)
		fmt.Println(*y) // y is 1
}