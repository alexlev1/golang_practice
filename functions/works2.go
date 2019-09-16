package main

import "fmt"

func checker(number uint) (int, bool) {
	if int(number) % 2 == 0 {
			return 1, true
	} else {
			return 0, false
	}
}

func main() {
		fmt.Println("Enter a number:")
		var number uint
		fmt.Scanf("%f", &number)
		fmt.Println(checker(number))
}