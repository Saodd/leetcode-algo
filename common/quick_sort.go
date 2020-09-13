package common

// 快速排序1：用栈实现的超复杂版本 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func QuickSortInt(li []int) {
	if len(li) == 0 {
		return
	}
	quickSortInt(li, 0, len(li)-1)
}

func quickSortInt(li []int, lo, hi int) {
	stack := quickSortStackInt{}
	stack.Push(lo, hi)
	for stack.Len() > 0 {
		x, y, _ := stack.Pop()
		mid := quickSortIntPartition(li, x, y)
		if mid-x > 15 {
			stack.Push(x, mid-1)
		} else if mid-x > 1 {
			quickSortIntSelectSort(li, x, mid-1)
		}
		if y-mid > 15 {
			stack.Push(mid+1, y)
		} else if y-mid > 1 {
			quickSortIntSelectSort(li, mid+1, y)
		}
	}
}

func quickSortIntPartition(li []int, lo, hi int) (mid int) {
	l, r := lo, hi
	//li[lo], li[(lo+hi)/2] = li[(lo+hi)/2], li[lo]
	midValue := li[lo]
	for l < r {
		for l <= hi { // 找大的
			if li[l] > midValue {
				break
			}
			l++
		}
		for r >= lo {
			if li[r] <= midValue {
				break
			}
			r--
		}
		if l < r {
			li[l], li[r] = li[r], li[l]
		} else {
			break
		}
	}
	li[lo], li[r] = li[r], li[lo]
	return r
}

func quickSortIntSelectSort(li []int, lo, hi int) {
	var min int
	for ; lo < hi; lo++ {
		min = lo
		for i := lo + 1; i <= hi; i++ {
			if li[i] < li[min] {
				min = i
			}
		}
		if lo != min {
			li[lo], li[min] = li[min], li[lo]
		}
	}
}

type quickSortStackInt struct {
	items []int
}

func (s *quickSortStackInt) Len() int {
	return len(s.items)
}

func (s *quickSortStackInt) Push(x, y int) {
	s.items = append(s.items, x, y)
}
func (s *quickSortStackInt) Pop() (int, int, error) {
	l := len(s.items)
	//if l < 2 {
	//	return 0, 0, errors.New("没有元素可以取出")
	//}
	x, y := s.items[l-2], s.items[l-1]
	s.items = s.items[:l-2]
	return x, y, nil
}

// 快速排序1 结束 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

// 快速排序2：不要用栈，只加了选择排序的基础版快排（反转排序） >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func QuickSortIntRev(li []int) {
	quickSortIntRev(li, 0, len(li)-1)
}

func quickSortIntRev(li []int, lo, hi int) {
	if hi-lo < 5 {
		quickSortIntSelectSortRev(li, lo, hi)
		return
	}
	mid := quickSortIntPartitionRev(li, lo, hi)
	quickSortIntRev(li, lo, mid-1)
	quickSortIntRev(li, mid+1, hi)
}

func quickSortIntPartitionRev(li []int, lo, hi int) (mid int) {
	l, r := lo, hi
	midValue := li[lo] // 比较处
	for l < r {
		for l <= hi {
			if li[l] < midValue { // 比较处
				break
			}
			l++
		}
		for r >= lo {
			if li[r] >= midValue { // 比较处
				break
			}
			r--
		}
		if l < r {
			li[l], li[r] = li[r], li[l]
		} else {
			break
		}
	}
	li[lo], li[r] = li[r], li[lo]
	return r
}

func quickSortIntSelectSortRev(li []int, lo, hi int) {
	var min int
	for ; lo < hi; lo++ {
		min = lo
		for i := lo + 1; i <= hi; i++ {
			if li[i] > li[min] { // 比较处
				min = i
			}
		}
		if lo != min {
			li[lo], li[min] = li[min], li[lo]
		}
	}
}

// 快速排序2 结束 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
