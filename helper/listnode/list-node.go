package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(values []int) *ListNode {
	var head = &ListNode{}
	var cur = head
	for _, value := range values {
		cur.Next = &ListNode{Val: value}
		cur = cur.Next
	}
	return head.Next
}

func Format(head *ListNode) []int {
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

// NewListWithCycle 负责创建一个带环的链表。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。
func NewListWithCycle(li []int, pos int) *ListNode {
	if pos >= len(li) {
		panic("错误的输入！pos应该小于li的长度！")
	}
	head := NewList(li)
	if pos >= 0 {
		var insectNode *ListNode
		var currentNode = &ListNode{Next: head}
		for p := range li {
			currentNode = currentNode.Next
			if p == pos {
				insectNode = currentNode
			}
		}
		currentNode.Next = insectNode
	}
	return head
}

func Walk(head *ListNode, step int) *ListNode {
	if head == nil {
		return nil
	}
	for i := 0; i < step && head.Next != nil; i++ {
		head = head.Next
	}
	return head
}
