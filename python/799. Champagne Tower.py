class Solution(object):
    def champagneTower(self, poured, query_row, query_glass):
        """
        :type poured: int
        :type query_row: int
        :type query_glass: int
        :rtype: float
        """
        dp = [[0]*100 for _ in range(100)]
        dp[0][0] = poured
        for i in range(1, 100):
            count = 0
            for j in range(0, i+1):
                ans = 0
                if j > 0:
                    ans += (dp[i-1][j-1]-1)/2.0 if dp[i-1][j-1] > 1 else 0
                if j < i:
                    ans += (dp[i-1][j]-1)/2.0 if dp[i-1][j] > 1 else 0
                dp[i][j] = ans
                if ans > 1: # 本层有杯子溢出
                    count += 1
            if count==0: # 本层没有杯子溢出，跳出循环
                break
        return min(1, dp[query_row][query_glass])
# 思路: dp
# time: O(N^2), space: O(N^2)