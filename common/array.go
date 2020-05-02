package common

// MinInt 找出两个整形中更小的值。
func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// IntArrayBinaryFind 数组中二分查找，返回目标值所在的索引，不存在则返回-1。
func IntArrayBinaryFind(list []int, target int) (index int) {
	var left, right = 0, len(list) - 1
	for right >= left {
		index = (left + right) / 2
		if target < list[index] {
			right = index - 1
		} else if target > list[index] {
			left = index + 1
		} else {
			return index
		}
	}
	return -1
}
