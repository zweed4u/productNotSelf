package main

import (
	"fmt"
)

func main() {
	array := []int{4, 5, 1, 8, 2}
	fmt.Println(array)
	sizeOfArray := len(array)
	// initialize two temp product arrays for each side
	leftProducts := []int{1}  // [1, 4, 20, 20, 160]
	rightProducts := []int{1} // [80, 16, 16, 2, 1]

	// loop through array regularly and multiply
	for index := range array {
		if index >= sizeOfArray-1 {
			break
		}
		leftProducts = append(leftProducts, array[index]*leftProducts[index])
	}
	fmt.Println(leftProducts)

	// loop through array backwards and multiply -
	// prepending to right products array
	for index := range array {
		if index >= sizeOfArray-1 {
			break
		}
		reverseIndex := len(array) - 1 - index // 4,3,2,1,0
		rightProductsReverseIndex := len(rightProducts) - 1 - index
		// composite literal to prepend - https://stackoverflow.com/a/53737602
		rightProducts = append([]int{array[reverseIndex] * rightProducts[rightProductsReverseIndex]}, rightProducts...)
	}
	fmt.Println(rightProducts)

	// go through temp arrays and multiply each
	sumArray := make([]int, len(leftProducts))
	for index := range leftProducts {
		sumArray[index] = leftProducts[index] * rightProducts[index]
	}
	fmt.Println(sumArray)
}
