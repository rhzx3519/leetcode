class Solution(object):
    def minOperationsMaxProfit(self, customers, boardingCost, runningCost):
        """
        :type customers: List[int]
        :type boardingCost: int
        :type runningCost: int
        :rtype: int
        """
        waiting_num = 0
        in_num = 0
        max_profit = 0
        rd = 0
        ans = -1
        while True:
            if rd < len(customers):
                waiting_num += customers[rd]
            in_num += min(4, waiting_num)
            waiting_num = max(0, waiting_num - 4)
            cur_profit = in_num*boardingCost - (rd+1)*runningCost
            if cur_profit > max_profit:
                max_profit = cur_profit
                ans = rd + 1
            rd += 1
            if rd >= len(customers) and waiting_num <= 0:
                break

        return ans