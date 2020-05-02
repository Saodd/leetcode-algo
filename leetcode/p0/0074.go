package p0

import "github.com/Saodd/leetcode-algo/common"

/*
74. 搜索二维矩阵

编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

示例 1:
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true

示例 2:
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func searchMatrix2(matrix [][]int, target int) bool {
	/*
		大致思路 ： （官方题解思路）二维坐标转换为一维坐标。
		执行用时 :	4 ms	, 在所有 Go 提交中击败了	98.17%	的用户
		内存消耗 :	3.8 MB	, 在所有 Go 提交中击败了	100.00%	的用户
	*/
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var y, x = len(matrix), len(matrix[0])
	var left, right = 0, x*y - 1
	for right >= left {
		mid := (left + right) / 2
		midValue := matrix[mid/x][mid%x]
		if target < midValue {
			right = mid - 1
		} else if target > midValue {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	/*
		大致思路 ： 先在纵向做二分，然后到横向做二分。
		执行用时 :	8 ms	, 在所有 Go 提交中击败了	79.58%	的用户
		内存消耗 :	3.8 MB	, 在所有 Go 提交中击败了	100.00%	的用户
	*/
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var xm, ym = len(matrix[0]) - 1, len(matrix) - 1
	if target < matrix[0][0] || target > matrix[ym][xm] {
		return false
	}
	// 先在纵向确定 target 在哪一行
	var up, row, down = 0, ym / 2, ym
	for {
		if target < matrix[row][0] {
			down = row
		} else if target > matrix[row][0] {
			if up == row {
				// 此时 down == up + 1, target 可能处在 down数组 上。
				// 无法像数组二分查找中那样直接+1，所以打个补丁额外判断一下。
				if target >= matrix[down][0] {
					row = down
				}
				break
			}
			up = row
		} else {
			return true
		}
		row = (up + down) / 2
	}
	// 再在一行中进行二分查找
	return common.IntArrayBinaryFind(matrix[row], target) >= 0
}
