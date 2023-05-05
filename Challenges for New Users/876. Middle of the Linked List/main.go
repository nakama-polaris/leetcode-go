package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func middleNode(head *ListNode) *ListNode {
	listNodeArray := make([]*ListNode, 0)

	// append head to listNodeArray while head is not nil
	for head != nil {
		listNodeArray = append(listNodeArray, head)
		head = head.Next
	}

	// return the middle node of listNodeArray
	return listNodeArray[len(listNodeArray)/2]
}
