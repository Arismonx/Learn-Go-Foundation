package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		fmt.Println("--------------------------")
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}

		}
	}
	return []int{}
}

func main() {
	num := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(num, target))
}
