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

// Gcd 负责计算最大公约数
func Gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return Gcd(y, tmp)
	} else {
		return y
	}
}
