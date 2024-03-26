package main

import "fmt"

func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for _, r := range nums {
		m[r]++
	}
	for k, v := range m {
		if v > 1 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}) == 2)
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}) == 3)
	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3}) == 3)
}
