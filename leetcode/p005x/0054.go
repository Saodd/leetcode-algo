package p005x

/*
54. 螺旋矩阵

给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]

示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var result = make([]int, len(matrix[0])*len(matrix))

	// 四个边界
	var right, down, left, up = len(matrix[0]) - 1, len(matrix) - 1, 0, 1
	// 当前指针的位置x,y
	var x, y = 0, 0
	// 循环填满结果数组
	for direct, p := 0, 0; p < len(result); p++ {
		result[p] = matrix[y][x]

		// 决定这次的前进方向
		switch direct {
		case 0: // 向右碰到边界
			if x == right {
				direct = 1
				right--
			}
		case 1: // 向下碰到边界
			if y == down {
				direct = 2
				down--
			}
		case 2: // 向左碰到边界
			if x == left {
				direct = 3
				left++
			}
		case 3: // 向上碰到边界
			if y == up {
				direct = 0
				up++
			}
		}
		// 前进一步
		switch direct {
		case 0:
			x++
		case 1:
			y++
		case 2:
			x--
		case 3:
			y--
		}
	}
	return result
}
