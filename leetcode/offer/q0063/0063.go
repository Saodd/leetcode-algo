/*
剑指 Offer 63. 股票的最大利润

假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

示例 2:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

限制：
0 <= 数组长度 <= 10^5

注意：本题与主站 121 题相同：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package q0063

/*
执行用时：4 ms, 在所有 Go 提交中击败了94.38%的用户
内存消耗：3 MB, 在所有 Go 提交中击败了67.23%的用户
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var minPrice, maxxProfit = prices[0], 0
	for _, price := range prices {
		profit := price - minPrice
		if profit > maxxProfit {
			maxxProfit = profit
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxxProfit
}
