/*
面试题 04.02. 最小高度树
难度：简单

给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-height-tree-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q1929

import "github.com/Saodd/leetcode-algo/helper/treenode"

type TreeNode = treenode.TreeNode

/*
2021-03-25
执行用时：4 ms, 在所有 Go 提交中击败了82.50%的用户
内存消耗：4.2 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func sortedArrayToBST(nums []int) *TreeNode {
	switch len(nums) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: nums[0]}
	default:
		mid := len(nums) / 2
		return &TreeNode{
			Val:   nums[mid],
			Left:  sortedArrayToBST(nums[:mid]),
			Right: sortedArrayToBST(nums[mid+1:]),
		}
	}
}
