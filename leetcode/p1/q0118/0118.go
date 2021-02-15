/*
118. 杨辉三角
难度：简单

给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	var rows = [][]int{{1}}
	for num := 2; num <= numRows; num++ {
		lastRow := rows[num-2]
		row := make([]int, num)
		row[0], row[num-1] = 1, 1
		for i := 1; i < num-1; i++ {
			row[i] = lastRow[i-1] + lastRow[i]
		}
		rows = append(rows, row)
	}
	return rows
}
