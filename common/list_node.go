package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsEqualListChain(a, b *ListNode) bool {
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

func CreateListChain(li []int) *ListNode {
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

func PrintListChain(x *ListNode) {
	count := 0
	for ; x != nil; x = x.Next {
		fmt.Printf("%d ", x.Val)
		count++
		if count == 100 {
			fmt.Print("...")
			break
		}
	}
	fmt.Println()
}

// CreateListChainCycle 负责创建一个带环的链表。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。
func CreateListChainCycle(li []int, pos int) *ListNode {
	if pos >= len(li) {
		panic("错误的输入！pos应该小于li的长度！")
	}
	head := CreateListChain(li)
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
