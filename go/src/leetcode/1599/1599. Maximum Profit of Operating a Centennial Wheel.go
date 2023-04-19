package main

/*
*
https://leetcode.cn/problems/maximum-profit-of-operating-a-centennial-wheel/description/
*/
func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) (ans int) {
	ans = -1
	var inNum, waitNum, maxProfit, curProfit int
	n := len(customers)
	for rd := 0; ; rd++ {
		if rd < n {
			waitNum += customers[rd]
		}
		inNum += min(4, waitNum)
		waitNum = max(0, waitNum-4)
		curProfit = inNum*boardingCost - runningCost*(rd+1)
		if curProfit > maxProfit {
			maxProfit = curProfit
			ans = rd + 1
		}
		if rd >= n && waitNum <= 0 {
			break
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
