/*
剑指 Offer 17. 打印从1到最大的n位数
难度：简单

输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]

说明：

用返回一个整数列表来代替打印
n 为正整数

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0017

// 这个简单解法没有考虑大数越界int32问题，但是也通过单元测试了，因为越界时的数组内存已经吓死人了。
// 我在做的时候，直觉就认为不可能会有越界问题，所以如果用这题来考察越界问题是不合理的。
func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}

	l := 1
	for i := 0; i < n; i++ {
		l *= 10
	}
	l -= 1

	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = i + 1
	}
	return res
}
