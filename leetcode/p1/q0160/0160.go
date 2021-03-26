/*
160. 相交链表
难度：简单（中等）

编写一个程序，找到两个单链表相交的起始节点。

注意：

如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0160

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
	"github.com/Saodd/leetcode-algo/leetcode/p1/q0142"
)

type ListNode = listnode.ListNode

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var a, b = headA, headB
	var goA, goB = true, true
	for goA || goB {
		if a.Next == nil {
			a = headB
			goA = false
		} else {
			a = a.Next
		}
		if b.Next == nil {
			b = headA
			goB = false
		} else {
			b = b.Next
		}
	}
	for a != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 找到A链表的尾部
	var tailA = headA
	for ; tailA.Next != nil; tailA = tailA.Next {
	}
	// 把A链表收尾相连做成环，最后要恢复原样
	tailA.Next = headA
	defer func() { tailA.Next = nil }()
	// 142题算法
	return q0142.DetectCycle(headB)
}
