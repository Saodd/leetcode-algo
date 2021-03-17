/*
剑指 Offer 60. n个骰子的点数

把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

示例 1:
输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]

示例2:
输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]

限制：
1 <= n <= 11

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0060

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了81.87%的用户
*/
func dicesProbability(n int) []float64 {
	res := make([]float64, 1)
	res[0] = 1.0
	for nn := 0; nn < n; nn++ {
		for i := range res {
			res[i] /= 6
		}
		newRes := make([]float64, len(res)+5)
		for i := 0; i < 6; i++ {
			for ii := range res {
				newRes[ii+i] += res[ii]
			}
		}
		res = newRes
	}
	return res
}
