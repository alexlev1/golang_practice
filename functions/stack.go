package main

import "fmt"

func f1() int {
	return f2()
}

func f2() (r int) {
	r = 5
	return
}

func f3() (int, int) {
	return 5, 6
}

func main() {
		fmt.Println(f1())
		x, y := f3()
		fmt.Println(x, y)
}

