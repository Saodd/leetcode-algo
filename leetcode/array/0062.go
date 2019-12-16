package array

import "github.com/Saodd/leetcode-algo/leetcode/maths"

/*
62. 不同路径

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右

示例 2:

输入: m = 7, n = 3
输出: 28

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func uniquePaths(m int, n int) int {
	m--
	n--
	// 保证n更小
	if n > m {
		n, m = m, n
	}
	// 准备分子和分母
	var numo, deno = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		numo[i] = m + i + 1
		deno[i] = i + 1
	}
	// 分子分母对消。最后答案一定是整数，分母一定能全部被分子消除掉。
	for _, d := range deno {
		// 把分子的元素集中一下，减少循环次数
		for numo[0] < 1<<25 && len(numo) > 1 {
			numo[0] = numo[0] * numo[1]
			numo[1], numo[len(numo)-1] = numo[len(numo)-1], numo[1]
			numo = numo[:len(numo)-1]
		}
		for i := range numo {
			// 分母元素能被分子整除就最好了，不能整除就去掉最大公约数
			if numo[i]%d == 0 {
				numo[i] /= d
				break
			} else {
				g := maths.Gcd(numo[i], d)
				numo[i] /= g
				d /= g
			}
		}
	}
	// 把分子剩下的元素乘起来
	var paths int = 1
	for _, n := range numo {
		paths *= n
	}
	return paths
}

// 另一种实现：对象式 ------------------------
func uniquePaths2(m int, n int) int {
	b := newUniquePathsBoard(m, n)
	return b.Paths(m-1, n-1)
}

type uniquePathsBoard [][]int

func newUniquePathsBoard(m, n int) uniquePathsBoard {
	var cache = make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	return cache
}

func (b uniquePathsBoard) Paths(m, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	if b[m][n] == 0 {
		b[m][n] = b.Paths(m-1, n) + b.Paths(m, n-1)
	}
	return b[m][n]
}

// 另一种实现：函数式 ------------------------
func uniquePaths1(m int, n int) int {
	// 创建一个二维数组
	var cache = make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	// 递归求解
	return uniquePathsRec(m-1, n-1, cache)
}

func uniquePathsRec(m, n int, cache [][]int) int {
	// 递归终止条件
	if m == 0 || n == 0 {
		return 1
	}
	// 查看之前是否计算过
	if cache[m][n] == 0 {
		cache[m][n] = uniquePathsRec(m-1, n, cache) + uniquePathsRec(m, n-1, cache)
	}
	return cache[m][n]
}
