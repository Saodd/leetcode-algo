// mySortInt 在这个包里练习常用的排序算法，附带测试用例。
package mySortInt

func SelectSort(li []int) {
	for i := range li {
		minI := i
		for ii := i; ii < len(li); ii++ {
			if li[ii] < li[minI] {
				minI = ii
			}
		}
		li[minI], li[i] = li[i], li[minI]
	}
}

func QuickSort(li []int) {
	if len(li) < 6 {
		SelectSort(li)
		return
	}

	midValue := li[0]
	var left, right = 1, len(li) - 1
	for left < right {
		for li[left] <= midValue && left < right {
			left++ // 最大值right, 不会越界
		}
		for li[right] >= midValue && right > 0 {
			right-- // 最小值0，不会越界
		}
		if left >= right {
			break
		}
		li[left], li[right] = li[right], li[left]
	}
	// 0号位置应该放小数，因此应该把right换过来
	li[0], li[right] = li[right], li[0]

	QuickSort(li[:right])
	QuickSort(li[right+1:])
}

func BubbleSort(li []int) {
	for max := len(li); max > 1; max-- {
		for i := 1; i < max; i++ {
			if li[i-1] > li[i] {
				li[i-1], li[i] = li[i], li[i-1]
			}
		}
	}
}
