package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var listNode0 ListNode
	var listNode1 ListNode
	var listNode2 ListNode

	listNode0.Val = 0
	listNode0.Next = &listNode1

	listNode1.Val = 2
	listNode1.Next = &listNode2

	listNode2.Val = 0

	head := listNode0

	fmt.Println(isPalindrome(&head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// array
	var array []int
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	// compare
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		if array[i] != array[j] {
			return false
		}
	}

	return true
}
