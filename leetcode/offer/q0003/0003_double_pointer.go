package q0003

/*
方案一：
使用快排思路，用双指针去遍历。
时间复杂度 NlogN ，空间复杂度 1 。
暴力计算的门槛值设置为8或者10时，成绩最好，40ms，击败95%
*/
func findRepeatNumber2(nums []int) int {
	if len(nums) < 6 {
		return findRepeatNumberSmall(nums)
	}
	mid, left, right := len(nums)/2, 1, len(nums)-1
	nums[0], nums[mid] = nums[mid], nums[0]
	bench := nums[0]
	for left < right {
		if nums[left] < bench {
			left++
			continue
		} else if nums[left] == bench {
			return bench
		}
		if nums[right] > bench {
			right--
			continue
		} else if nums[right] == bench {
			return bench
		}

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	if nums[0] > nums[right] {
		nums[0], nums[right] = nums[right], nums[0]
	}
	if num := findRepeatNumber2(nums[:right]); num != -1 {
		return num
	}
	if num := findRepeatNumber2(nums[right:]); num != -1 {
		return num
	}
	return -1
}

func findRepeatNumberSmall(nums []int) int {
	for i, v := range nums {
		for _, vv := range nums[i+1:] {
			if v == vv {
				return v
			}
		}
	}
	return -1 // 因为限制了数字最小是0
}
