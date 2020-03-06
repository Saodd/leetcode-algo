package p0

import "sort"

/*
47. 全排列 II

给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)
	g0047.len = len(nums)
	g0047.nums = nums
	g0047.temp = make([]int, len(nums))
	g0047.used = make([]bool, len(nums))
	g0047.result = make([][]int, 0, 6)

	recurPermuteUnique(0)
	return g0047.result
}

var g0047 = struct {
	nums   []int
	temp   []int
	len    int
	result [][]int
	used   []bool
}{}

func recurPermuteUnique(pos int) {
	if pos == g0047.len {
		newSolution := make([]int, len(g0047.temp))
		copy(newSolution, g0047.temp)
		g0047.result = append(g0047.result, newSolution)
		return
	}
	for i := 0; i < g0047.len; i++ {
		if !g0047.used[i] {
			if i > 0 && !g0047.used[i-1] && g0047.nums[i] == g0047.nums[i-1] {
				continue
			}
			g0047.temp[pos] = g0047.nums[i]
			g0047.used[i] = true
			recurPermuteUnique(pos + 1)
			g0047.used[i] = false
		}
	}
}

// // ---------------------------------------
//func permuteUnique(nums []int) [][]int {
//    if len(nums)==0{
//        return [][]int{}
//    }
//
//    len = len(nums) - 1
//    temp = nums
//    sort.Ints(temp)
//
//    result = make([][]int, 0, 5)
//
//    recurPermuteUnique(0)
//    return result
//}
//
//var temp []int
//var len int
//var result [][]int
//
//func recurPermuteUnique(pos int) {
//    if pos == len {
//        newSolution := make([]int, len(temp))
//        copy(newSolution, temp)
//        result = append(result, newSolution)
//        return
//    }
//
//    recurPermuteUnique(pos + 1)
//    for i := pos + 1; i <= len; i++ {
//        for j:=pos; j<i; j++{
//            if temp[j]==temp[i]{
//                goto NEXTI
//            }
//        }
//        temp[pos], temp[i] = temp[i], temp[pos]
//        recurPermuteUnique(pos + 1)
//        temp[pos], temp[i] = temp[i], temp[pos]
//        NEXTI:
//    }
//}
