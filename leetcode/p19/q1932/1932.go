/*
面试题 04.05. 合法二叉搜索树

实现一个函数，检查一棵二叉树是否为二叉搜索树。

示例1:
输入:
    2
   / \
  1   3
输出: true

示例2:
输入:
    5
   / \
  1   4
    / \
   3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
    根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/legal-binary-search-tree-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q1932

import (
	"github.com/Saodd/leetcode-algo/helper/treenode"
)

type TreeNode = treenode.TreeNode

/*
2021-03-25
执行用时：4 ms, 在所有 Go 提交中击败了97.97%的用户
内存消耗：5.1 MB, 在所有 Go 提交中击败了88.51%的用户
注意：子树也要保持父节点的大小要求
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isValidBSTRecur(root, false, false, 0, 0)
}

func isValidBSTRecur(root *TreeNode, f1, f2 bool, v1, v2 int) bool {
	// 要求 root!=nil
	if f1 && root.Val <= v1 {
		return false
	}
	if f2 && root.Val >= v2 {
		return false
	}
	if root.Left != nil {
		if ok := isValidBSTRecur(root.Left, f1, true, v1, root.Val); !ok {
			return false
		}
	}
	if root.Right != nil {
		if ok := isValidBSTRecur(root.Right, true, f2, root.Val, v2); !ok {
			return false
		}
	}
	return true
}
