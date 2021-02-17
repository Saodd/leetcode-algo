package q0046

import (
	"strconv"
)

/*
剑指 Offer 46. 把数字翻译成字符串

给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

提示：

0 <= num < 231

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	// x, y, z 分别代表 后2位、后1位、当前位的答案
	var x, y, z = 0, 0, 1
	for i := 0; i < len(numStr); i++ {
		// 滚动一次，当前位置在最坏情况下与上一位保持相同数量
		x, y, z = y, z, z

		// 防御性代码，防止截取字符串截到-1位置
		if i == 0 {
			continue
		}

		// 截取字符串并且判断条件
		prv := numStr[i-1 : i+1]
		if prv <= "25" && prv >= "10" {
			z = y + x // 如果可以组合，则要加上后2位的答案
		} else {
			z = y
		}
	}
	return z
}

func translateNum1(num int) int {
	if num < 0 {
		return 0
	}
	if num < 10 {
		return 1
	}

	var digs []int // len>=2
	for ; num > 0; num /= 10 {
		digs = append(digs, num%10)
	}

	var results = make([]int, len(digs))
	results[0] = 1
	f1 := func(i int) { // i>=1
		results[i] = results[i-1]
	}
	f2 := func(i int) { // i>=1
		if i == 1 {
			results[i] = results[i-1] + 1
		} else {
			results[i] = results[i-1] + results[i-2]
		}
	}

	for i := 1; i < len(digs); i++ {
		shi, ge := digs[i], digs[i-1]
		if shi > 2 {
			f1(i)
		} else if shi == 2 {
			if ge > 5 {
				f1(i)
			} else {
				f2(i)
			}
		} else if shi == 1 {
			f2(i)
		} else {
			// f==0
			f1(i)
		}
	}

	return results[len(results)-1]
}
