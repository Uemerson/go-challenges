package main

import "fmt"

func firstUniqChar(s string) int {
	var arr [26]int // arr := [26]int{}
	for _, r := range s {
		arr[r-97]++ //arr[r-'a']++
	}
	for i, r := range s {
		if arr[r-97] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode") == 0)
	fmt.Println(firstUniqChar("loveleetcode") == 2)
	fmt.Println(firstUniqChar("aabb") == -1)
}
