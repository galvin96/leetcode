package linkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}

	list := l3
	plus := 0

	for l1 != nil || l2 != nil || plus != 0 {
		sum := getNodeVal(l1) + getNodeVal(l2) + plus

		node := &ListNode{
			Val: sum % 10,
		}
		list.Next = node
		list = list.Next

		plus = sum / 10

		l1 = getNodeNext(l1)
		l2 = getNodeNext(l2)
	}
	return l3.Next
}

func (list *ListNode) List() {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func linkedList() {
	l1 := ListNode{
		Val: 9,
	}
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)

	l2 := ListNode{
		Val: 9,
	}
	l2.Insert(9)
	l2.Insert(9)
	l2.Insert(9)

	addTwoNumbers(&l1, &l2)
}
