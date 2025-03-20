package main

import (
	"fmt"
)

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0 // Handle empty array case
	}
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0 // Handle empty array case
	}
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println("Hello World")
	arr := []int{84, 81, 100, 87}
	fmt.Println(findMin(arr))
	fmt.Println(findMax(arr))
}
