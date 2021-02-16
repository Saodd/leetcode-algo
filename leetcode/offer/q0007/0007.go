/*
剑指 Offer 07. 重建二叉树
难度：中等

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

限制：

0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0007

/*
执行用时：4 ms, 在所有 Go 提交中击败了95.89%的用户
内存消耗：4.2 MB, 在所有 Go 提交中击败了38.28%的用户

在定位 rootPos 的时候是用线性扫描，这种方式的时间复杂度较高（应该是 NlogN）。
一种解决方案是对inorder扫描一次，建立哈希表。这样用额外的 N 的空间复杂度，换取了较低的时间复杂度（N）。
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	var rootVal = preorder[0]
	var rootPos int
	for inorder[rootPos] != rootVal {
		rootPos++
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:rootPos+1], inorder[:rootPos]),
		Right: buildTree(preorder[rootPos+1:], inorder[rootPos+1:]),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
