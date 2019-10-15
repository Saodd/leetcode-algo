package p005x

/*
56. 合并区间

给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	QuickSortInterval(intervals)

	var merges = make([][]int, 0)
	var merge = []int{intervals[0][0], intervals[0][1]}
	for i := range intervals {
		left, right := intervals[i][0], intervals[i][1]
		if left <= merge[1] {
			if right > merge[1] {
				merge[1] = right
			}
		} else {
			merges = append(merges, merge)
			merge = []int{left, right}
		}
	}
	merges = append(merges, merge)
	return merges
}

// 快速排序 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func QuickSortInterval(li [][]int) {
	if len(li) == 0 {
		return
	}
	quickSortInterval(li, 0, len(li)-1)
}

func quickSortInterval(li [][]int, lo, hi int) {
	stack := quickSortStackInterval{}
	stack.Push(lo, hi)
	for stack.Len() > 0 {
		x, y, _ := stack.Pop()
		mid := quickSortIntervalPartition(li, x, y)
		if mid-x > 15 {
			stack.Push(x, mid-1)
		} else if mid-x > 1 {
			quickSortIntervalSelectSort(li, x, mid-1)
		}
		if y-mid > 15 {
			stack.Push(mid+1, y)
		} else if y-mid > 1 {
			quickSortIntervalSelectSort(li, mid+1, y)
		}
	}
}

func quickSortIntervalPartition(li [][]int, lo, hi int) (mid int) {
	l, r := lo, hi
	//li[lo], li[(lo+hi)/2] = li[(lo+hi)/2], li[lo]
	midValue := li[lo][0]
	for l < r {
		for l <= hi { // 找大的
			if li[l][0] > midValue {
				break
			}
			l++
		}
		for r >= lo {
			if li[r][0] <= midValue {
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

func quickSortIntervalSelectSort(li [][]int, lo, hi int) {
	var min int
	for ; lo < hi; lo++ {
		min = lo
		for i := lo + 1; i <= hi; i++ {
			if li[i][0] < li[min][0] {
				min = i
			}
		}
		if lo != min {
			li[lo], li[min] = li[min], li[lo]
		}
	}
}

type quickSortStackInterval struct {
	items []int
}

func (s *quickSortStackInterval) Len() int {
	return len(s.items)
}

func (s *quickSortStackInterval) Push(x, y int) {
	s.items = append(s.items, x, y)
}
func (s *quickSortStackInterval) Pop() (int, int, error) {
	l := len(s.items)
	//if l < 2 {
	//	return 0, 0, errors.New("没有元素可以取出")
	//}
	x, y := s.items[l-2], s.items[l-1]
	s.items = s.items[:l-2]
	return x, y, nil
}

// 快速排序 结束 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
