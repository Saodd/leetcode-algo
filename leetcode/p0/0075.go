package p0

/*
75. 颜色分类

给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]

进阶：

一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sortColors(nums []int) {
	/*
		大致思路： 从左向右扫描，遇到0就丢到左边，遇到2就丢到右边。一共标记三个指针（左、中、右）。
		执行用时 :	0 ms	, 在所有 Go 提交中击败了	100.00%	的用户
		内存消耗 :	2 MB	, 在所有 Go 提交中击败了	100.00%	的用户
	*/
	if len(nums) <= 1 {
		return
	}

	// left 是 下一个放置0的位置。right是下一个2的位置。
	var left, right = 0, len(nums) - 1
	for left <= right && nums[left] == 0 {
		left++
	}
	for left <= right && nums[right] == 2 {
		right--
	}

	// 用now作为当前索引，遍历一次
	for now := left; now <= right; {
		switch nums[now] {
		case 0:
			if now != left {
				nums[now], nums[left] = nums[left], nums[now]
			}
			left++
			// left 指向的是非0索引，理论上可能是1或2。但此时由于是从左向右遍历，因此左边不可能剩下2。
			// 因此left只会指向1，因此 now 可以放心前进一步。
			now++
		case 1:
			now++
		case 2:
			nums[now], nums[right] = nums[right], nums[now]
			right--
			// right 指向的是非2索引，理论上可能是1或0。
			// 因此 now 不能前进，等下一个循环再判断一次。
		}
	}
}
