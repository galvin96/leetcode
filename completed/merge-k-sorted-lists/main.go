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

type Tree struct {
	Val   int
	left  *Tree
	right *Tree
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

func insertTree(tree *Tree, newVal int) {
	clone := tree
	for {
		if newVal < clone.Val {
			if clone.left == nil {
				clone.left = &Tree{
					Val: newVal,
				}
				return
			} else {
				clone = clone.left
			}
		} else {
			if clone.right == nil {
				clone.right = &Tree{
					Val: newVal,
				}
				return
			} else {
				clone = clone.right
			}
		}
	}
}

func listTree(tree *Tree) []int {
	if tree == nil {
		return []int{}
	}
	ret := append(listTree(tree.left), tree.Val)
	ret = append(ret, listTree(tree.right)...)
	return ret
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ret ListNode

	/**
	* Find the min node
	**/
	var tree *Tree = nil
	for _, node := range lists {
		for node != nil {
			if tree == nil {
				tree = &Tree{
					Val: node.Val,
				}
			} else {
				insertTree(tree, node.Val)
			}
			node = node.Next
		}
	}

	list := listTree(tree)
	for _, val := range list {
		insert(&ret, val)
	}

	return ret.Next
}

func main() {
	l1 := ListNode{
		Val: 1,
	}
	insert(&l1, 4)
	insert(&l1, 5)

	l2 := ListNode{
		Val: 1,
	}
	insert(&l2, 3)
	insert(&l2, 4)

	l3 := ListNode{
		Val: 2,
	}
	insert(&l3, 6)

	mergeKLists([]*ListNode{&l1, &l2, &l3})
}
