package q0031

import (
	"sort"
)

/*
2021-03-23
执行用时：4 ms, 在所有 Go 提交中击败了74.04%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了97.03%的用户
*/
func nextPermutation3(nums []int) {
	// 两个数的平均值，避免整形溢出
	half := func(left, right int) int {
		return (right-left)/2 + left
	}

	// 传入一个递减排序的数组
	findMinGreatPos := func(nums []int, compare int) int {
		var left, right = 0, len(nums)
		var mid int
		for mid = half(left, right); left < right-1; mid = half(left, right) {
			if nums[mid] <= compare {
				right = mid
			} else {
				left = mid
			}
		}
		return left
	}

	// 降序数组排序，直接用交换排序
	swapSort := func(nums []int) {
		maxI := len(nums) / 2
		for i := 0; i < maxI; i++ {
			j := len(nums) - 1 - i
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 1. 找到需要进位的位置（即从右向左数，第一个升序排列的位置）
	var downPos int
	for downPos = len(nums) - 2; downPos >= 0; downPos-- {
		if nums[downPos] < nums[downPos+1] {
			break
		}
	}
	if downPos < 0 {
		sort.Ints(nums)
		return
	}
	// 2. 从后面降序排列数组中，找到下一个值（即最小的比前面那个数字大的数字）
	descendingList := nums[downPos+1:]
	minGreatPos := findMinGreatPos(descendingList, nums[downPos])
	// 3. 交换数字并排序
	nums[downPos], descendingList[minGreatPos] = descendingList[minGreatPos], nums[downPos]
	swapSort(descendingList)
}
