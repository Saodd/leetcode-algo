package p005x

/*
51. N皇后

n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，
该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例:

输入: 4
输出: [
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]

解释: 4 皇后问题存在两个不同的解法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func solveNQueens(n int) [][]string {
	// 初始化各项全局变量
	len0051 = n
	temp0051 = make([]int, n)
	pattern0051 = make([]byte, n)
	for i := range temp0051 {
		temp0051[i] = i
		pattern0051[i] = '.'
	}
	result0051 = [][]string{}
	// 开始递归
	recSolveNQueens(0)

	return result0051
}

var len0051 int           // 皇后的数量
var temp0051 []int        // 当前棋盘的摆放情况
var result0051 [][]string // 最后返回的结果
var pattern0051 []byte    // 帮助生成字符串

func recSolveNQueens(n int) {
	var num int
	// 递归到最后一位，如果满足斜线条件就记录这个答案
	if n == len0051-1 {
		num = temp0051[n]
		for j := 0; j < n; j++ {
			if temp0051[j]+n-j == num || temp0051[j]-n+j == num {
				return
			}
		}
		recordNQueens()
		return
	}
	// 不是最后一位，就回溯+递归
	for i := n; i < len0051; i++ {
		if i == n {
			num = temp0051[n]
			for j := 0; j < n; j++ {
				if temp0051[j]+n-j == num || temp0051[j]-n+j == num {
					goto NEXTNUM
				}
			}
			recSolveNQueens(n + 1)
		} else {
			temp0051[n], temp0051[i] = temp0051[i], temp0051[n]
			num = temp0051[n]
			for j := 0; j < n; j++ {
				if temp0051[j]+n-j == num || temp0051[j]-n+j == num {
					temp0051[n], temp0051[i] = temp0051[i], temp0051[n]
					goto NEXTNUM
				}
			}
			recSolveNQueens(n + 1)
			temp0051[n], temp0051[i] = temp0051[i], temp0051[n]
		}
	NEXTNUM:
	}
}

func recordNQueens() {
	// 记录当前的答案
	solution := make([]string, len0051)
	for i, n := range temp0051 {
		pattern0051[i] = 'Q'
		solution[n] = string(pattern0051)
		pattern0051[i] = '.'
	}
	result0051 = append(result0051, solution)
}
