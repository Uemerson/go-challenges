package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var head *ListNode
	var node *ListNode
	index := 0
	for curr := list1; curr != nil; {
		if index < a || index > b {
			if head != nil {
				tmp := node
				node = &ListNode{
					Val:  curr.Val,
					Next: nil,
				}
				tmp.Next = node
			}
			if head == nil {
				node = &ListNode{
					Val:  curr.Val,
					Next: nil,
				}
				head = node
			}
		}
		if index == a {
			for r := list2; r != nil; {
				if head != nil {
					tmp := node
					node = &ListNode{
						Val:  r.Val,
						Next: nil,
					}
					tmp.Next = node
				}
				if head == nil {
					node = &ListNode{
						Val:  curr.Val,
						Next: nil,
					}
					head = node
				}
				r = r.Next
			}
		}
		index++
		curr = curr.Next
	}
	return head
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

func showListNode(node *ListNode) {
	curr := node
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func compareListNode(node *ListNode, list []int) bool {
	curr := node
	index := 0
	for curr != nil {
		if index >= len(list) || curr.Val != list[index] {
			return false
		}
		index++
		curr = curr.Next
	}
	return true
}

func main() {
	node1 := newListNode([]int{10, 1, 13, 6, 9, 5})
	node2 := newListNode([]int{1000000, 1000001, 1000002})
	node3 := mergeInBetween(node1, 3, 4, node2)
	// showListNode(node3)

	node4 := newListNode([]int{0, 1, 2, 3, 4, 5, 6})
	node5 := newListNode([]int{1000000, 1000001, 1000002, 1000003, 1000004})
	node6 := mergeInBetween(node4, 2, 5, node5)
	// showListNode(node6)

	fmt.Println(compareListNode(node3, []int{10, 1, 13, 1000000, 1000001, 1000002, 5}) == true)
	fmt.Println(compareListNode(node6, []int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}) == true)
}
