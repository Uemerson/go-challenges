package main

import "fmt"

func isPalindrome(x int) bool {
	n := x
	if n < 0 {
		return false
	}

	reversed := 0

	for n != 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}

	if reversed == x {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPalindrome(121) == true)
	fmt.Println(isPalindrome(-121) == false)
	fmt.Println(isPalindrome(10) == false)
}
