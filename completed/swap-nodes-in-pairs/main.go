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

func swapPairs(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil && curr.Next != nil {
		t1 := curr
		t2 := curr.Next
		t3 := curr.Next.Next

		t1.Next = t3
		t2.Next = t1
		if curr == head {
			head = t2
		} else {
			prev.Next = t2
		}
		prev = t1
		curr = t3
	}

	return head
}

func main() {
	l1 := ListNode{
		Val: 1,
	}
	insert(&l1, 2)
	insert(&l1, 3)
	insert(&l1, 4)
	insert(&l1, 5)
	swapPairs(&l1)

	l2 := ListNode{
		Val: 1,
	}
	swapPairs(&l2)

	swapPairs(nil)
}
