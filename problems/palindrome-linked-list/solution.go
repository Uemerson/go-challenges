package main

import "fmt"

func isPalindrome(head *ListNode) bool {
	var arr []int
	for curr := head; curr != nil; {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	check := len(arr) - 1
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[check] {
			return false
		}
		check--
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(values []int) *ListNode {
	var curr *ListNode
	var head *ListNode
	for _, r := range values {
		tmp := ListNode{
			Val:  r,
			Next: nil,
		}
		if curr == nil {
			curr = &tmp
			head = curr
			continue
		}
		curr.Next = &tmp
		curr = &tmp
	}
	return head
}

func main() {
	fmt.Println(isPalindrome(newListNode([]int{1, 2, 2, 1})) == true)
	fmt.Println(isPalindrome(newListNode([]int{1, 2})) == false)
}
