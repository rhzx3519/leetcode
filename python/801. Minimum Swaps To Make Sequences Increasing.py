class Solution(object):
    def minSwap(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        n = len(A)
        dp = [[0, 0] for _ in range(n)]
        dp[0][0] = 0
        dp[0][1] = 1
        for i in range(1, n):
            if A[i-1] < A[i] and B[i-1] < B[i]:
                if A[i-1] < B[i] and B[i-1] < A[i]: # 任意交换或者交换，取最优解
                    dp[i][0] = min(dp[i-1][0], dp[i-1][1])
                    dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
                else:
                    dp[i][0] = dp[i-1][0] # i不交换，上一步也不能交换
                    dp[i][1] = dp[i-1][1] + 1 # i交换，上一步也要交换
            else:
                dp[i][0] = dp[i-1][1] # i不交换，上一步必交换
                dp[i][1] = dp[i-1][0] + 1 # i交换，上一步不能交换
        return min(dp[n-1][0], dp[n-1][1])

# https://leetcode-cn.com/problems/minimum-swaps-to-make-sequences-increasing/solution/dong-tai-gui-hua-de-yi-ban-bu-zou-by-li-zhi-chao-4/