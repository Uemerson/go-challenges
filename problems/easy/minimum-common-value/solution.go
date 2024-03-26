package main

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	m := make(map[int]int)
	for _, r := range nums1 {
		m[r]++
	}
	for _, r := range nums2 {
		_, ok := m[r]
		if ok {
			return r
		}
	}
	return -1
}

func main() {
	fmt.Println(getCommon([]int{1, 2, 3}, []int{2, 4}) == 2)
	fmt.Println(getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}) == 2)
	fmt.Println(getCommon([]int{1000000000, 1000000000}, []int{1000000000}) == 1000000000)
}
