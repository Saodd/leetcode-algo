/*
141. 环形链表
难度：简单

给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例2：

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0141

import (
	"github.com/Saodd/leetcode-algo/helper/listnode"
)

type ListNode = listnode.ListNode

/*
2021-03-27（难度：简单）
执行用时：8 ms, 在所有 Go 提交中击败了82.99%的用户
内存消耗：4.3 MB, 在所有 Go 提交中击败了24.45%的用户
*/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	var slow, fast = head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
