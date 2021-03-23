package q0031

func nextPermutation2(nums []int) {
	// 二分查找
	indexOfMinGreater := func(nums []int, target int) int {
		var gt, lt int
		for gt, lt = 0, len(nums); lt-gt > 1; {
			mid := (gt + lt) / 2
			if nums[mid] > target {
				gt = mid
			} else {
				lt = mid
			}
		}
		return gt
	}

	// 交换排序
	swapSortInt := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	if len(nums) < 2 {
		return
	}
	var top int
	for top = len(nums) - 1; top > 0; top-- {
		if nums[top-1] < nums[top] {
			break
		}
	}
	if top == 0 {
		swapSortInt(nums)
		return
	}
	exchId := indexOfMinGreater(nums[top:], nums[top-1]) + top
	nums[top-1], nums[exchId] = nums[exchId], nums[top-1]
	swapSortInt(nums[top:])
	return
}

var NextPermutation = nextPermutation2
