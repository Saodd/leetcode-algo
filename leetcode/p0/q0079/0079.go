package q0079

/*
79. 单词搜索

给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func exist(board [][]byte, word string) bool {
	Word = []byte(word)
	Board = board
	Depth = len(word)
	Path = make([][]bool, 0, len(board))
	for range board {
		Path = append(Path, make([]bool, len(board[0])))
	}
	first := Word[0]
	for x, row := range board {
		for y, b := range row {
			if b == first {
				Path[x][y] = true
				if Next(x, y, 1) {
					return true
				}
				Path[x][y] = false
			}
		}
	}
	return false
}

var (
	Board [][]byte
	Word  []byte
	Depth int
	Path  [][]bool
)

func Next(x, y int, depth int) bool {
	if depth == Depth {
		return true
	}
	first := Word[depth]
	xmax, ymax := len(Board), len(Board[0])
	if x := x + 1; x < xmax && Board[x][y] == first {
		if !Path[x][y] {
			Path[x][y] = true
			if Next(x, y, depth+1) {
				return true
			}
			Path[x][y] = false
		}
	}
	if y := y + 1; y < ymax && Board[x][y] == first {
		if !Path[x][y] {
			Path[x][y] = true
			if Next(x, y, depth+1) {
				return true
			}
			Path[x][y] = false
		}
	}
	if x := x - 1; x >= 0 && Board[x][y] == first {
		if !Path[x][y] {
			Path[x][y] = true
			if Next(x, y, depth+1) {
				return true
			}
			Path[x][y] = false
		}
	}
	if y := y - 1; y >= 0 && Board[x][y] == first {
		if !Path[x][y] {
			Path[x][y] = true
			if Next(x, y, depth+1) {
				return true
			}
			Path[x][y] = false
		}
	}
	return false
}

/*
日期：2021-02-13
时间：实现+优化 约1小时

感想：
其实优化点在于如何记录已经搜索过的路径。
第一次用list存+map检验，超过时间限制了；
第二次用map存，时间排名16%；
最后才想到用board相同大小的bool二维数组来实现，时间排名95%，达到最优解。
*/
