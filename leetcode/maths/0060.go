package maths

/*
60. 第k个排列

给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。

示例 1:

输入: n = 3, k = 3
输出: "213"

示例 2:

输入: n = 4, k = 9
输出: "2314"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getPermutation(n int, k int) string {
	// 先把所有元素生成
	var nums = make([]byte, n)
	for i := byte(0); i < byte(n); i++ {
		nums[i] = i + '1'
	}
	// 因为n, k 都是从1开始计算，因此把他们都减掉1
	n--
	k--
	// 逐位做除法，找出k落在哪一段
	var fac = Factorial(n - 1)
	for i := 0; i < n; i++ {
		upInsertBytes(i, i+k/fac, nums)
		k %= fac
		fac /= (n - i)
	}
	return string(nums)
}

// upInsertBytes 把ori位置的元素前移到tar位置，中间的元素相应后移一位。
func upInsertBytes(tar, ori int, bs []byte) {
	temp := bs[ori]
	for i := ori; i > tar; i-- {
		bs[i] = bs[i-1]
	}
	bs[tar] = temp
}
