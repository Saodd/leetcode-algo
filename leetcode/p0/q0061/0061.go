/*
61. 旋转链表
难度：中等（简单）

给定一个链表，旋转链表，将链表每个节点向右移动k个位置，其中k是非负数。

示例1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步:0->1->2->NULL
向右旋转 4 步:2->0->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0061

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
)

type ListNode = listnode.ListNode

/*
2021-03-27
执行用时：4 ms, 在所有 Go 提交中击败了68.15%的用户
内存消耗：2.6 MB, 在所有 Go 提交中击败了94.81%的用户
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var tail = head
	var count = 1
	for tail.Next != nil {
		count++
		tail = tail.Next
	}
	tail.Next = head

	var move = count - k%count
	for i := 0; i < move; i++ {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil
	return head
}
