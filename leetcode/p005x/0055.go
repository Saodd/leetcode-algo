package p005x

/*
55. 跳跃游戏

给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。

示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ，
所以你永远不可能到达最后一个位置。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	var endPoint = len(nums)
	var touchPoint = nums[0]
	for pos := 0; pos < touchPoint; pos++ {
		// 当前点能够触及的最远点
		newPoint := nums[pos] + pos + 1
		// 如果是比记录中更远的点，那就记录下来
		if newPoint > touchPoint {
			// 如果达到终点，就返回
			if newPoint >= endPoint {
				return true
			}
			touchPoint = newPoint
		}
	}
	return false
}
