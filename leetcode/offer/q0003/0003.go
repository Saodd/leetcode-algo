/*
剑指 Offer 03. 数组中重复的数字
难度：简单

找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

限制：

2 <= n <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0003

/*
方案三：
考虑到一个特殊条件：nums里的数字都是符合下标条件的，所以可以利用下标来做。
用一个指针从左向右遍历。把每个数字都放到这个数字作为下标的位置上去，如果发现有两个数字放在一起了，说明重复了。
时间复杂度：N，空间复杂度：1，理论最优。
但是实际表现并没有明显优势，用时36ms，击败96%
*/
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		num := nums[i]
		if i == num { // 因为置换过来的数字依然大概率不是符合位置的，因此可能需要在一个位置上循环判定。
			i++
			continue
		}
		if nums[num] == num {
			return num
		}
		nums[i], nums[num] = nums[num], nums[i]
	}
	return -1
}
