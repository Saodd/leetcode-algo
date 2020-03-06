package p0

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
	g0051.len = n
	g0051.temp = make([]int, n)
	g0051.pattern = make([]byte, n)
	for i := range g0051.temp {
		g0051.temp[i] = i
		g0051.pattern[i] = '.'
	}
	g0051.result = [][]string{}
	// 开始递归
	recSolveNQueens(0)

	return g0051.result
}

var g0051 = struct {
	len     int        // 皇后的数量
	temp    []int      // 当前棋盘的摆放情况
	result  [][]string // 最后返回的结果
	pattern []byte     // 帮助生成字符串
}{}

func recSolveNQueens(n int) {
	var num int
	// 递归到最后一位，如果满足斜线条件就记录这个答案
	if n == g0051.len-1 {
		num = g0051.temp[n]
		for j := 0; j < n; j++ {
			if g0051.temp[j]+n-j == num || g0051.temp[j]-n+j == num {
				return
			}
		}
		recordNQueens()
		return
	}
	// 不是最后一位，就回溯+递归
	for i := n; i < g0051.len; i++ {
		if i == n {
			num = g0051.temp[n]
			for j := 0; j < n; j++ {
				if g0051.temp[j]+n-j == num || g0051.temp[j]-n+j == num {
					goto NEXTNUM
				}
			}
			recSolveNQueens(n + 1)
		} else {
			g0051.temp[n], g0051.temp[i] = g0051.temp[i], g0051.temp[n]
			num = g0051.temp[n]
			for j := 0; j < n; j++ {
				if g0051.temp[j]+n-j == num || g0051.temp[j]-n+j == num {
					g0051.temp[n], g0051.temp[i] = g0051.temp[i], g0051.temp[n]
					goto NEXTNUM
				}
			}
			recSolveNQueens(n + 1)
			g0051.temp[n], g0051.temp[i] = g0051.temp[i], g0051.temp[n]
		}
	NEXTNUM:
	}
}

func recordNQueens() {
	// 记录当前的答案
	solution := make([]string, g0051.len)
	for i, n := range g0051.temp {
		g0051.pattern[i] = 'Q'
		solution[n] = string(g0051.pattern)
		g0051.pattern[i] = '.'
	}
	g0051.result = append(g0051.result, solution)
}
