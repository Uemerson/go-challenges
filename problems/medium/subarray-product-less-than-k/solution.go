package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	count := 0
	curr := 0
	next := 0
	tmp := 0
	for curr < len(nums) {
		if next == len(nums) {
			curr++
			next = curr
			continue
		}
		if curr == next && nums[curr] < k {
			next++
			count++
			tmp = nums[curr]
			continue
		}
		if curr == next && nums[curr] >= k {
			curr++
			next = curr
			continue
		}
		if tmp*nums[next] < k {
			tmp *= nums[next]
			next++
			count++
			continue
		}
		if tmp*nums[next] >= k {
			curr++
			next = curr
			continue
		}
	}
	return count
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100) == 8)
	fmt.Println(numSubarrayProductLessThanK([]int{1, 1, 1}, 1) == 0)
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0) == 0)
}
