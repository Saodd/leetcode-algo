package maths

// Factorial 负责计算n的阶乘
func Factorial(n int) int {
	if n < 1 {
		return 1
	}
	var result int = 1
	for ; n > 1; n-- {
		result *= n
	}
	return result
}
