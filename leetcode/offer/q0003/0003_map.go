package q0003

/*
方案二：
把数组里的值全部丢进map里去检验。
时间复杂度 N，空间复杂度 N 。
时间排名40%，空间26%，很差。
*/
func findRepeatNumber3(nums []int) int {
	book := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := book[num]; exist {
			return num
		}
		book[num] = struct{}{}
	}
	return -1
}
