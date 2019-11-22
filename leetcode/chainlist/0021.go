package chainlist

import . "github.com/Saodd/leetcode-algo/common"

/*
将两个有序链表合并为一个新的有序链表并返回。
新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	}
	if l1 == nil {
		return l2
	}
	var head *ListNode = &ListNode{}
	var tail *ListNode = head
	for {
		if l1.Val < l2.Val {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
			if l1 == nil {
				tail.Next = l2
				break
			}
		} else {
			tail.Next = l2
			tail = tail.Next
			l2 = l2.Next
			if l2 == nil {
				tail.Next = l1
				break
			}
		}
	}
	return head.Next
}
