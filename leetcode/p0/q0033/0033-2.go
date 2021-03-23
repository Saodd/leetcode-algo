package q0033

/*
2021-03-23 太难了，重点做
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了57.98%的用户
*/
func search2(nums []int, target int) int {
	if len(nums) < 2 {
		for i, n := range nums {
			if n == target {
				return i
			}
		}
		return -1
	}

	var half = func(a, b int) int { return (b-a)/2 + a }

	var max = len(nums) - 1
	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := half(left, right)
		if target == nums[mid] {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[max] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
