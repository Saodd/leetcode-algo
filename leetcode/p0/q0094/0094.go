/*
94. 二叉树的中序遍历
难度：中等

给定一个二叉树的根节点 root ，返回它的 中序 遍历。

提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
package q0094

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderTraversalRun(root, &res)
	return res
}

func inorderTraversalRun(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversalRun(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalRun(root.Right, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
感想：这也忒简单了……3分钟写完一次通过。中等界的耻辱。
看了下其他人的解答，还有基于栈的实现，的确是一种思路。我之前好像做过（看到栈就想起来了），所以不再重复了。
*/
