package main

import "fmt"

func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}

	curr := 0
	for i := 1; i <= n; i++ {
		curr += i
		next := 0
		for j := i; j <= n; j++ {
			next += j
		}
		if curr == next {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotInteger(8) == 6)
	fmt.Println(pivotInteger(1) == 1)
	fmt.Println(pivotInteger(4) == -1)
}
