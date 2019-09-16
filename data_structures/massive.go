package main

import "fmt"

func main() {
		var x [5]float64 // Array
		x[0] = 98
		x[1] = 93
		x[2] = 77
		x[3] = 82
		x[4] = 83

		var total float64 = 0
		for i := 0; i < len(x); i++ {
				total += x[i]
		}
		fmt.Println(total / float64(len(x)))

		// Another code with range
		array := [5]float64{ 98, 93, 77, 82, 83 }
		var totalSecond float64 = 0
		for _, value := range array {
				totalSecond += value
		}
		fmt.Println(totalSecond / float64(len(array)))
}