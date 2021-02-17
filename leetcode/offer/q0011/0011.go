/*
剑指 Offer 11. 旋转数组的最小数字
难度：简单

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：
输入：[3,4,5,1,2]
输出：1

示例 2：
输入：[2,2,2,0,1]
输出：0

注意：本题与主站 154 题相同：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0011

/*
评论：
跟官方题解略有不同。在 y==z 的情况下，官方题解的意思是忽略掉right位置，即right--然后继续循环。
而我的思路，在 y==z 的情况下，我把数组从mid位置切开，对两边分别递归然后返回更小的那个。
代码有可以精简的部分，但是影响不大，所以依然保留最易读的样子。

时间复杂度：最坏 (N)，平均 (logN)
空间复杂度：最坏 (logN)，因为要递归。正常 (1)

执行用时：4 ms, 在所有 Go 提交中击败了91.12%的用户
内存消耗：3.1 MB, 在所有 Go 提交中击败了59.08%的用户
*/
func minArray(numbers []int) int {
	var left, right = 0, len(numbers) - 1
	for right-left > 8 {
		mid := right/2 + left/2
		var x, y, z = numbers[left], numbers[mid], numbers[right]
		if x < y {
			left = mid
		} else if x > y {
			right = mid
		} else { // x==y
			if y > z {
				left = mid
			} else { // x==y==z
				return minInt(minArray(numbers[1:mid]), minArray(numbers[mid+1:]))
			}
		}
	}
	return minArraySmall(numbers[left : right+1])
}

func minArraySmall(numbers []int) int {
	if len(numbers) == 0 {
		return 1<<32 - 1
	}
	var min = numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
