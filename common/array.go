package common

// MinInt 找出两个整形中更小的值。
func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
