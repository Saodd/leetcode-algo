/*
剑指 Offer 14- I. 剪绳子
难度：中等

给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36

提示：
2 <= n <= 58

注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jian-sheng-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0014_1

/*
评论：
这其实是一道数学题。
由乘法分配律可知，和一定的两个数相乘时，两个数越接近，乘积越大。同理可推导：和一定的N个数相乘时，N个数越接近，乘积越大。
因此，我们尝试由2段、3段……推导到n/2段，每种情况令每一段长度尽可能相等，然后记下最大的乘积即可。

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了77.44%的用户
*/
func cuttingRope(n int) int {
	secMax := (n + 1) / 2
	productMax := 1
	for sec := 2; sec <= secMax; sec++ { // 2段、3段……n/2段的情况分别求解
		secLength := n / sec
		secExtra := n % sec
		product := 1
		// 让每段长度尽可能相等，即一部分长度是 secLength，另一部分长度是 secLength+1
		for i := 0; i < secExtra; i++ {
			product *= secLength + 1
		}
		for i := 0; i < (sec - secExtra); i++ {
			product *= secLength
		}
		// 记录最大值
		if product > productMax {
			productMax = product
		}
	}
	return productMax
}
