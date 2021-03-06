package p0

import . "github.com/Saodd/leetcode-algo/common"

/*
92. 反转链表 II

反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var dumHead = &ListNode{Next: head}
	var start = dumHead
	for i := 1; i < m; i++ {
		start = start.Next
	}

	var a, b, c *ListNode
	a = start // start, a, b, c 都不会是nil，因为m<n<=链表长度
	b = a.Next
	c = b.Next
	for i := m; i < n; i++ {
		a, b, c = b, c, c.Next
		b.Next = a
	}
	start.Next, start.Next.Next = b, c
	return dumHead.Next
}
