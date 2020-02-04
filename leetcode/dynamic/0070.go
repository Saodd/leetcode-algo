package dynamic

/*
70. 爬楼梯

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	a, b := 2, 3
	for i := 3; i < n; i++ {
		b, a = b+a, b
	}
	return b
	/*
		第三次：
		执行用时 :	0 ms	, 在所有 Go 提交中击败了	100.00%	的用户
		内存消耗 :	1.9 MB	, 在所有 Go 提交中击败了	87.35%	的用户

		注：
		题解中还给出了利用矩阵乘法的算法，复杂度是log(n)。
		由于矩阵乘法是纯粹的数学（线性代数），已经不是编程算法思想了，所以暂时不考虑。（主要还是我忘记矩阵乘法怎么做了……）
	*/
}

func climbStairs2(n int) int {
	if n <= 3 {
		return n
	}
	dyn := make([]int, n)
	dyn[0], dyn[1], dyn[2] = 1, 2, 3
	for i := 3; i < n; i++ {
		dyn[i] = dyn[i-1] + dyn[i-2]
	}
	return dyn[n-1]
	/*
		第二次
		执行用时 :	0 ms	, 在所有 Go 提交中击败了	100.00%	的用户
		内存消耗 :	1.9 MB	, 在所有 Go 提交中击败了	42.24%	的用户
	*/
}

func climbStairs1(n int) int {
	if n <= 3 {
		return n
	}
	dyn := make([]int, n)
	dyn[0], dyn[1], dyn[2] = 1, 2, 3
	return climbStairs1Dyn(n-1, dyn)
	/*
		第一次：
		执行用时 :	0 ms	, 在所有 Go 提交中击败了	100.00%	的用户
		内存消耗 :	2.1 MB	, 在所有 Go 提交中击败了	5.89%	的用户
	*/
}

func climbStairs1Dyn(n int, dyn []int) int {
	t := dyn[n]
	if t == 0 {
		t = climbStairs1Dyn(n-1, dyn) + climbStairs1Dyn(n-2, dyn)
		dyn[n] = t
	}
	return t
}
