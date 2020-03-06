package p1

import (
	. "github.com/Saodd/leetcode-algo/common"
)

/*
148. 排序链表

在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4

示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func sortList(head *ListNode) *ListNode {
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
	head2 := sortList(slow.Next)
	slow.Next = nil
	head = sortList(head)
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
