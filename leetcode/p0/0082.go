package p0

import . "github.com/Saodd/leetcode-algo/common"

/*
82. 删除排序链表中的重复元素 II

给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func deleteDuplicates(head *ListNode) *ListNode {
	left := &ListNode{}
	newHead := left
	for right := head; right != nil; right = right.Next {
		dupFlag := false
		for right.Next != nil && right.Val == right.Next.Val {
			dupFlag = true
			right = right.Next
		}
		if !dupFlag {
			left.Next = right
			left = right
		}
	}
	left.Next = nil
	return newHead.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	var dupMap = map[int]bool{}
	// 用map记录重复值
	for p := head; p != nil; p = p.Next {
		_, ok := dupMap[p.Val]
		if ok {
			dupMap[p.Val] = true
		} else {
			dupMap[p.Val] = false
		}
	}
	// 删除重复节点
	left := &ListNode{}
	newHead := left
	for right := head; right != nil; right = right.Next {
		if !dupMap[right.Val] {
			left.Next = right
			left = right
		}
	}
	left.Next = nil
	return newHead.Next
}
