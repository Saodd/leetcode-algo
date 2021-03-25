package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func Unmarshal(values []int) *ListNode {
	var head = &ListNode{}
	var cur = head
	for _, value := range values {
		cur.Next = &ListNode{Val: value}
		cur = cur.Next
	}
	return head.Next
}

func Marshal(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var res = []int{head.Val}
	var cur = head.Next
	for cur != nil && cur != head { // 防止循环链表导致死循环
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func Concat(front, back *ListNode) *ListNode {
	if front == nil {
		return back
	}
	var tail = front
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = back
	return front
}
