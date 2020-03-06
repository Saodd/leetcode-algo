package p0

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"

示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

提交成绩：
执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗 :2.2 MB, 在所有 Go 提交中击败了82.95%的用户
*/
func addBinary(a string, b string) string {
	// 令a的长度大于等于b
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	// 准备变量
	var temp []byte = make([]byte, len(a)+1)
	copy(temp[1:], a)
	ex := false
	x, y := len(temp)-1, len(b)-1
	// 把b加入temp中
	for y >= 0 {
		if ex {
			temp[x] += (b[y] - '0' + 1)
		} else {
			temp[x] += (b[y] - '0')
		}
		if temp[x] > '1' {
			temp[x] -= 2
			ex = true
		} else {
			ex = false
		}
		x--
		y--
	}
	// temp剩下的部分要继续处理进位
	if ex {
		for ; x > 0; x-- {
			if temp[x] == '0' {
				temp[x] = '1'
				ex = false
				break
			} else {
				temp[x] = '0'
			}
		}
	}
	// 整理最后结果返回
	if ex {
		temp[0] = '1'
		return string(temp)
	}
	return string(temp[1:])
}
