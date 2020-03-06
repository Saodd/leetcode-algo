package p0

import "github.com/Saodd/leetcode-algo/common"

/*
64. 最小路径和

给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			grid[x][y] += minPathSumValue(x, y, grid)
		}
	}
	return grid[m-1][n-1]
}

func minPathSumValue(x, y int, grid [][]int) int {
	if x == 0 && y == 0 {
		return 0
	}
	if x == 0 {
		return grid[0][y-1]
	}
	if y == 0 {
		return grid[x-1][0]
	}
	return common.MinInt(grid[x-1][y], grid[x][y-1])
}

// 第二次提交：在第一次的基础上稍作优化 ----------------------------------------
func minPathSum2(grid [][]int) int {
	maxX, maxY := len(grid), len(grid[0])
	{
		for i := 1; i < maxY; i++ {
			grid[0][i] += grid[0][i-1]
		}
		for i := 1; i < maxX; i++ {
			grid[i][0] += grid[i-1][0]
		}
	}
	var minLevel int = maxX
	if maxX > maxY {
		minLevel = maxY
	}
	for level := 1; level < minLevel; level++ {
		for i := level; i < maxY; i++ {
			if grid[level][i-1] > grid[level-1][i] {
				grid[level][i] += grid[level-1][i]
			} else {
				grid[level][i] += grid[level][i-1]
			}
		}
		for i := level + 1; i < maxX; i++ {
			if grid[i][level-1] > grid[i-1][level] {
				grid[i][level] += grid[i-1][level]
			} else {
				grid[i][level] += grid[i][level-1]
			}
		}
	}
	return grid[maxX-1][maxY-1]
	/*
	   执行用时 :8 ms, 在所有 golang 提交中击败了91.42%的用户
	   内存消耗 :3.9 MB, 在所有 golang 提交中击败了92.42%的用户
	*/
}

// 第一次提交：建立一个空的二维数组，分别对行、列计算最小路径 --------------------------------------
func minPathSum1(grid [][]int) int {
	maxX, maxY := len(grid), len(grid[0])
	paths := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		paths[i] = make([]int, maxY)
	}
	paths[0][0] = grid[0][0]
	x, y := 0, 0
	for x < maxX && y < maxY {
		if x == 0 {
			for i := y + 1; i < maxY; i++ {
				paths[x][i] = paths[x][i-1] + grid[x][i]
			}
			for i := x + 1; i < maxX; i++ {
				paths[i][y] = paths[i-1][y] + grid[i][y]
			}
		} else {
			for i := y; i < maxY; i++ {
				if paths[x][i-1] > paths[x-1][i] {
					paths[x][i] = paths[x-1][i] + grid[x][i]
				} else {
					paths[x][i] = paths[x][i-1] + grid[x][i]
				}
			}
			for i := x; i < maxX; i++ {
				if paths[i][y-1] > paths[i-1][y] {
					paths[i][y] = paths[i-1][y] + grid[i][y]
				} else {
					paths[i][y] = paths[i][y-1] + grid[i][y]
				}
			}
		}
		x++
		y++
	}
	return paths[maxX-1][maxY-1]
}
