package best_time_to_buy_and_sell_stock

import (
	"fmt"
	"testing"
)

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
//
// 注意：你不能在买入股票前卖出股票。
//
//  
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 示例 2:
//
// 输入: [7,6,4,3,1]
// 输出: 0
// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

func Test(t *testing.T) {
	prices := []int{7,4,2,10,3,1,5,8}
	res := maxProfit(prices)
	fmt.Println(res)
	t.Log(res)

}

// BF
// func maxProfit(prices []int) int {
// 	pLen := len(prices) - 1
// 	i := 0
// 	profit := 0
// 	// day := 0
// 	for ; i < pLen; i++ {
// 		for j := i + 1; j <= pLen; j++ {
// 			curProfit := prices[j] - prices[i]
// 			if curProfit > profit {
// 				// day = j + 1
// 				profit = curProfit
// 			}
// 		}
// 	}
//
// 	return profit
// }

// 滑动窗口
func maxProfit(prices []int) int {
	pLen := len(prices) - 1
	iEnd := pLen - 1
	i, j := 0, 1
	profit := 0
	for i <= iEnd && j <= pLen {
		if prices[i] >= prices[j]  {
			i = j
			j++
		} else {
			curProfit := prices[j] - prices[i]
			if curProfit > profit {
				profit = curProfit
			}
			j++
		}
	}

	return profit
}
