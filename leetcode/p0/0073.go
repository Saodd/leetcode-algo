package p0

/*
73. 矩阵置零

给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

示例 1:
输入:
[ [1,1,1],
  [1,0,1],
  [1,1,1]]
输出:
[ [1,0,1],
  [0,0,0],
  [1,0,1]]

示例 2:
输入:
[ [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]]
输出:
[ [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]]

进阶:

一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/set-matrix-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func setZeroes(matrix [][]int) {
	/*
		执行用时 :	16 ms	, 在所有 Go 提交中击败了	82.06%	的用户
		内存消耗 :	6 MB	, 在所有 Go 提交中击败了	100.00%	的用户
	*/
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	var xl, yl = len(matrix[0]), len(matrix)

	// 因为后面要用第一行和第一列做标记，所以提前判断一下
	var xClear, yClear bool
	for x := 0; x < xl; x++ {
		if matrix[0][x] == 0 {
			xClear = true
			break
		}
	}
	if xClear {
		defer func() {
			for x := 0; x < xl; x++ {
				matrix[0][x] = 0
			}
		}()
	}
	for y := 0; y < yl; y++ {
		if matrix[y][0] == 0 {
			yClear = true
			break
		}
	}
	if yClear {
		defer func() {
			for y := 0; y < yl; y++ {
				matrix[y][0] = 0
			}
		}()
	}

	// 在第一行和第一列打上标记
	for y := 1; y < yl; y++ {
		for x := 1; x < xl; x++ {
			if matrix[y][x] == 0 {
				matrix[y][0] = 0
				matrix[0][x] = 0
			}
		}
	}

	// 执行标记
	for x := 1; x < xl; x++ {
		// 这里根据第一行的0标记来清除整列
		if matrix[0][x] == 0 {
			for y := 1; y < yl; y++ {
				matrix[y][x] = 0
			}
		}
	}
	for y := 1; y < yl; y++ {
		// 这里根据第一列的0标记来清除整行
		if matrix[y][0] == 0 {
			for x := 1; x < xl; x++ {
				matrix[y][x] = 0
			}
		}
	}
}
