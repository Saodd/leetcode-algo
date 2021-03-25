package p1

import "github.com/Saodd/leetcode-algo/helper/treenode"

/*
109. 有序链表转换二叉搜索树

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	// 用快慢指针找到中点，并在中点前后将链表切分成两部分
	var before, slow, fast *ListNode = nil, head, head.Next
	for fast != nil && fast.Next != nil {
		before = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 递归
	if before != nil {
		before.Next = nil
		return &TreeNode{Val: slow.Val, Left: sortedListToBST(head), Right: sortedListToBST(slow.Next)}
	} else {
		return &TreeNode{Val: slow.Val, Right: sortedListToBST(slow.Next)}
	}
}

type TreeNode = treenode.TreeNode
