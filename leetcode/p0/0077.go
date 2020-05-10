package p0

/*
77. 组合

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combine(n int, k int) [][]int {
	/*
		大致思路： (回朔法)小数放前面，大数放后面。用迭代的思路，下一层迭代只能使用比上一层更大的数字。
		执行用时 :	12 ms	, 在所有 Go 提交中击败了	76.79%	的用户
		内存消耗 :	6.3 MB	, 在所有 Go 提交中击败了	100.00%	的用户

		其他思路： 动态规划法。用公式表达 combine(n,k) = combine(n,k-1)+combine(n-1,k-1)。缺点是空间复杂度爆炸。
		官方题解： 字典序 (二进制排序) 组合
	*/
	var mod = make([]int, k)
	var result [][]int
	return combineRecur(result, mod, 0, 1, n)
}

func combineRecur(result [][]int, mod []int, pos, left, right int) [][]int {
	if pos == len(mod)-1 {
		for i := left; i <= right; i++ {
			mod[pos] = i
			var newMod = make([]int, len(mod))
			copy(newMod, mod)
			result = append(result, newMod)
		}
	} else {
		for i := left; i <= right; i++ {
			mod[pos] = i
			result = combineRecur(result, mod, pos+1, i+1, right)
		}
	}
	return result
}
