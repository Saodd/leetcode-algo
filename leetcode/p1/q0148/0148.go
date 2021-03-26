/*
148. 排序链表
难度：中等

在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4

示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0148

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
)

type ListNode = listnode.ListNode

/*
2021-03-27（难度：中等偏上）
执行用时：1480 ms, 在所有 Go 提交中击败了5.08%的用户
内存消耗：8.2 MB, 在所有 Go 提交中击败了18.35%的用户
分析：用的快排思路，看来并不适合链表，链表的特性适合使用归并排序（合并有序链表）。
*/
func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	{
		var slow, fast = head, head.Next
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		if slow.Next != nil {
			temp := head
			head = slow.Next
			slow.Next = slow.Next.Next
			head.Next = temp
		}
	}

	var value = head.Val
	var small, big, smallHead, bigHead *ListNode
	var cur = head.Next
	for cur != nil {
		if cur.Val <= value {
			if small == nil {
				small = cur
				smallHead = small
			} else {
				small.Next = cur
				small = small.Next
			}
		} else {
			if big == nil {
				big = cur
				bigHead = big
			} else {
				big.Next = cur
				big = big.Next
			}
		}
		cur = cur.Next
	}

	if small != nil {
		small.Next = nil
	}
	if big != nil {
		big.Next = nil
	}
	head.Next = nil
	smallHead = sortList2(smallHead)
	bigHead = sortList2(bigHead)

	head.Next = bigHead
	if smallHead == nil {
		return head
	}
	cur = smallHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return smallHead
}

func sortList1(head *ListNode) *ListNode {
	// 划分终点
	if head == nil || head.Next == nil {
		return head
	}
	// 把链表对半划分
	var slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 递归，分别对两半链表排序
	head2 := sortList1(slow.Next)
	slow.Next = nil
	head = sortList1(head)
	// 合并两个有序链表
	if head2 == nil {
		return head
	}
	var newHead *ListNode
	if head.Val <= head2.Val {
		newHead = head
		head = head.Next
		if head == nil {
			newHead.Next = head2
			return newHead
		}
	} else {
		newHead = head2
		head2 = head2.Next
		if head2 == nil {
			newHead.Next = head
			return newHead
		}
	}
	var p = newHead
	for {
		if head.Val <= head2.Val {
			p.Next = head
			p = head
			head = head.Next
			if head == nil {
				p.Next = head2
				break
			}
		} else {
			p.Next = head2
			p = head2
			head2 = head2.Next
			if head2 == nil {
				p.Next = head
				break
			}
		}
	}
	return newHead
}
