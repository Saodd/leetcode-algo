/*
93. 复原IP地址
难度：中等

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。

示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

示例 2：

输入：s = "0000"
输出：["0.0.0.0"]

示例 3：

输入：s = "1111"
输出：["1.1.1.1"]

示例 4：

输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]

示例 5：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

提示：

0 <= s.length <= 3000  (Lewin注：这个条件应该是错的)
s 仅由数字组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0093

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	ctx := MyContext{
		originS: s,
		Res:     nil,
		temp:    make([]string, 4),
	}
	ctx.Run(s, 0)
	return ctx.Res
}

type MyContext struct {
	originS string
	Res     []string
	temp    []string
}

func (c *MyContext) Run(s string, depth int) {
	if depth == 4 {
		newRes := strings.Join(c.temp, ".")
		if len(newRes) == len(c.originS)+3 {
			c.Res = append(c.Res, newRes)
		}
		return
	}
	for i := 0; i < 3 && i < len(s); i++ {
		sub := s[:i+1]
		if !isValid(sub) {
			return
		}
		c.temp[depth] = sub
		c.Run(s[i+1:], depth+1)
	}
}

func isValid(s string) bool {
	switch a := s[0]; a {
	case '0':
		return len(s) == 1
	case '1':
		return true
	case '2':
		if len(s) == 3 {
			return s[1:] <= "55"
		}
		return true
	default:
		return len(s) < 3
	}
}
