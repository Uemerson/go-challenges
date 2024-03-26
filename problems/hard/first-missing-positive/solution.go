package main

import "fmt"

func firstMissingPositive(nums []int) int {
	m := make([]int, len(nums)+1)

	for _, n := range nums {
		if n > 0 && n <= len(nums) {
			m[n]++
		}
	}

	for i := 1; i < len(nums)+1; i++ {
		if m[i] == 0 {
			return i
		}
	}

	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}) == 3)
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}) == 2)
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}) == 1)
}
