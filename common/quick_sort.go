package common

// 快速排序 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
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

// 快速排序 结束 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
