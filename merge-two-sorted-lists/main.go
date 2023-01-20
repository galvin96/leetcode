package main

import "fmt"

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

func getNodeVal(l *ListNode) int {
	if l != nil {
		return l.Val
	} else {
		return 0
	}
}

func getNodeNext(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	} else {
		return nil
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = ListNode{}

	for list1 != nil || list2 != nil {
		if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
			insert(&head, list1.Val)
			list1 = list1.Next
		} else {
			insert(&head, list2.Val)
			list2 = list2.Next
		}
	}

	return head.Next
}

func main() {
	l1 := ListNode{
		Val: 1,
	}
	insert(&l1, 2)
	insert(&l1, 4)

	l2 := ListNode{
		Val: 1,
	}
	insert(&l2, 3)
	insert(&l2, 4)

	mergeTwoLists(&l1, &l2)
}
