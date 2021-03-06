package p0

import . "github.com/Saodd/leetcode-algo/common"

/*
86. 分隔链表

给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func partition(head *ListNode, x int) *ListNode {
	var smallHead, bigHead = &ListNode{}, &ListNode{}
	var smallTail, bigTail = smallHead, bigHead
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			smallTail.Next = p
			smallTail = p
		} else {
			bigTail.Next = p
			bigTail = p
		}
	}
	bigTail.Next = nil
	smallTail.Next = bigHead.Next
	return smallHead.Next
}
