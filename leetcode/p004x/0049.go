package p004x

import (
	"unsafe"
)

/*
49. 字母异位词分组

给定一个字符串数组，将字母异位词组合在一起。
字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func groupAnagrams(strs []string) [][]string {
	var dictionary = map[string][]string{}

	var temp string
	for i := range strs {
		temp = countAlpha(strs[i])
		if value, ok := dictionary[temp]; ok {
			dictionary[temp] = append(value, strs[i])
		} else {
			dictionary[temp] = []string{strs[i]}
		}
	}

	var result [][]string
	for _, v := range dictionary {
		result = append(result, v)
	}

	return result
}

var counter = make([]byte, 26)

func countAlpha(s string) (count string) {
	for i := 0; i < 26; i++ {
		counter[i] = 0
	}
	for _, i := range s {
		counter[i-'a']++
	}
	return string(counter)
}

// 方案1 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func groupAnagrams1(strs []string) [][]string {
	var dictionary = map[string][]string{}

	var p []byte
	var temp string
	for i := range strs {
		p = []byte(strs[i])
		QuickSortByte(p)
		temp = *(*string)(unsafe.Pointer(&p))
		if value, ok := dictionary[temp]; ok {
			dictionary[temp] = append(value, strs[i])
		} else {
			dictionary[temp] = []string{strs[i]}
		}
	}

	var result [][]string
	for _, v := range dictionary {
		result = append(result, v)
	}

	return result
}

// 辅助 >>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func QuickSortByte(li []byte) {
	if len(li) == 0 {
		return
	}
	quickSortByte(li, 0, len(li)-1)
}

func quickSortByte(li []byte, lo, hi int) {
	stack := quickSortStack{}
	stack.Push(lo, hi)
	for stack.Len() > 0 {
		x, y, _ := stack.Pop()
		mid := quickSortBytePartition(li, x, y)
		if mid-x > 15 {
			stack.Push(x, mid-1)
		} else if mid-x > 1 {
			quickSortByteSelectSort(li, x, mid-1)
		}
		if y-mid > 15 {
			stack.Push(mid+1, y)
		} else if y-mid > 1 {
			quickSortByteSelectSort(li, mid+1, y)
		}
	}
}

func quickSortBytePartition(li []byte, lo, hi int) (mid int) {
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

func quickSortByteSelectSort(li []byte, lo, hi int) {
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

type quickSortStack struct {
	items []int
}

func (s *quickSortStack) Len() int {
	return len(s.items)
}

func (s *quickSortStack) Push(x, y int) {
	s.items = append(s.items, x, y)
}
func (s *quickSortStack) Pop() (int, int, error) {
	l := len(s.items)
	//if l < 2 {
	//	return 0, 0, errors.New("没有元素可以取出")
	//}
	x, y := s.items[l-2], s.items[l-1]
	s.items = s.items[:l-2]
	return x, y, nil
}
