package main

import "fmt"

// Сложнааа!!! Чего здесь происходит???
func makeEvenGenerator() func() uint {
		i := uint(0)
		return func() (ret uint) {
				ret = i
				i += 2
				return
		}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}