package main

import "fmt"

func main() {
		sliceNew := make([]float64, 4, 10) // Slice create
		sliceNew[0] = 2
		sliceNew[1] = 3
		fmt.Println(sliceNew)

		slice := make([]float64, 4)
		fmt.Println(slice)

		// Slice declare
		sliceDeclare := []int{1,2,3,4}
		fmt.Println(sliceDeclare)

		arr := [5]float64{1,2,3,4,5}
		x := arr[1:3] // Make slice from array [low:high]
		fmt.Println(x)

		// Append for slice
		slice1 := []int{1,2,3}
		slice2 := append(slice1, 4, 5)
		fmt.Println(slice1, slice2)

		// Copy for slice
		slice3 := []int{1,2,3}
		slice4 := make([]int, 2)
		copy(slice4, slice3)
		fmt.Println(slice3, slice4)
}