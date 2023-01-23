package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) List() {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func insert(list *ListNode, newVal int) {
	for list.Next != nil {
		list = list.Next
	}

	var newNode = ListNode{
		Val: newVal,
	}
	list.Next = &newNode
}

func length(list *ListNode) int {
	length := 0
	for list != nil {
		length++
		list = list.Next
	}
	return length
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head

	len := length(head)

	/**
	* remove the first node
	**/
	if len-n == 0 {
		return head.Next
	}

	for i := 1; i < len-n; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next

	return head
}

func main() {
	// l1 := ListNode{
	// 	Val: 1,
	// }
	// insert(&l1, 2)
	// insert(&l1, 3)
	// insert(&l1, 4)
	// insert(&l1, 5)
	// removeNthFromEnd(&l1, 2)

	l2 := ListNode{
		Val: 1,
	}
	insert(&l2, 2)
	removeNthFromEnd(&l2, 2)
}
