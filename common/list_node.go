package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsEqualListInt(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	// 必须同时为nil
	if a != b {
		return false
	}
	return true
}

func CreateListInt(li []int) *ListNode {
	if li == nil || len(li) == 0 {
		return nil
	}
	var head = &ListNode{}
	var tail = head
	for _, val := range li {
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}
	return head.Next
}

func PrintListInt(x *ListNode) {
	for x != nil {
		fmt.Print(x.Val, ' ')
	}
	fmt.Print('\n')
}
