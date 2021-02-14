/*
95. 不同的二叉搜索树 II
难度：中等

给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。

示例：

输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0095

func generateTrees(n int) []*TreeNode {
	nums := make([]int, n)
	for i, _ := range nums {
		nums[i] = i + 1
	}
	nodes := Run(nums)
	return nodes
}

func Run(nums []int) (nodes []*TreeNode) {
	if len(nums) == 0 {
		return []*TreeNode{nil}
	}
	if len(nums) == 1 {
		return []*TreeNode{{Val: nums[0]}}
	}
	for i, num := range nums {
		left, right := nums[:i], nums[i+1:]
		leftNodes, rightNodes := Run(left), Run(right)
		for _, l := range leftNodes {
			for _, r := range rightNodes {
				nodes = append(nodes, &TreeNode{
					Val:   num,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
