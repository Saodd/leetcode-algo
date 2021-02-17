/*
剑指 Offer 13. 机器人的运动范围
难度：中等

地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。
请问该机器人能够到达多少个格子？

示例 1：
输入：m = 2, n = 3, k = 1
输出：3

示例 2：
输入：m = 3, n = 1, k = 0
输出：1

提示：
1 <= n,m <= 100
0 <= k<= 20

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0013

/*
评论：
一开始以为会有投机取巧的算法，想了半天，发现还是需要逐个格子遍历。
主要思路是先判断格子坐标是否合格，然后从(0,0)开始传播可达性。最后同时符合坐标合法和可达性的点就是符合的点，计数就可以了。
虽然对比官方题解稍微啰嗦了一点，但是复杂度是一样的。

官方题解最大的优化是，搜索方向只向下和向右。而我的是朝四个方向搜索（因为没有考虑到特性，实际上也是可以只向两个方向搜索的）。

空间复杂度：(mn)
时间复杂度：(mn)

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了66.16%的用户
*/
func movingCount(m int, n int, k int) int {
	// 1. 初始化格子
	var board = make([][]int, m)
	for i := range board {
		board[i] = make([]int, n)
	}
	// 2. 判断合法
	for x, row := range board {
		kk := k - x/10 - x%10
		if kk < 0 {
			continue
		}
		for y := range row {
			if kk >= y/10+y%10 {
				board[x][y] = 1
			}
		}
	}
	// 3. 判断可到达
	propagate(board, 0, 0)
	// 4. 计数
	count := 0
	for _, row := range board {
		for _, num := range row {
			if num&0b11 == 0b11 {
				count++
			}
		}
	}
	return count
}

func propagate(board [][]int, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	if board[x][y]&0b10 > 0 || board[x][y]&0b1 == 0 {
		return
	}
	board[x][y] |= 0b10
	propagate(board, x+1, y)
	//propagate(board, x-1, y)
	propagate(board, x, y+1)
	//propagate(board, x, y-1)
}
