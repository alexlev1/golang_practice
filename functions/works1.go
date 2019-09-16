package main

import "fmt"

func sum(slice []float64) float64 {
		total := 0.0
		for _, v := range slice {
				total += v
		}
		return total
}

func main() {
		slice := []float64{1,2,3}
		fmt.Println(sum(slice))
}