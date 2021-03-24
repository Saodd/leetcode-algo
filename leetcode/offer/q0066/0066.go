/*
剑指 Offer 66. 构建乘积数组

给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。

不能使用除法。

示例:
输入: [1,2,3,4,5]
输出: [120,60,40,30,24]

提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0066

/*
执行结果：通过
执行用时：32 ms, 在所有 Go 提交中击败了42.70%的用户
内存消耗：9.3 MB, 在所有 Go 提交中击败了15.30%的用户
*/
func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	b := make([]int, len(a))
	for i := range b {
		b[i] = 1
	}

	// 从左向右
	var product int = 1
	for i := 1; i < len(b); i++ {
		product *= a[i-1]
		b[i] = product
	}

	// 从右向左
	product = 1
	for i := len(b) - 2; i >= 0; i-- {
		product *= a[i+1]
		b[i] *= product
	}

	return b
}
