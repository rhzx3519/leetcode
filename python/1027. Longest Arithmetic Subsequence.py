class Solution(object):
    def longestArithSeqLength(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        n = len(A)
        dp = collections.defaultdict(dict)
        ans = 1
        for i in range(1, n):
            for j in range(i):
                diff = A[i] - A[j]
                if diff not in dp[j]:
                    dp[i][diff] = 2
                else:
                    dp[i][diff] = dp[j].get(diff, 0) + 1
                ans = max(ans, dp[i][diff])
        return ans

# 思路：动态规划，记录dp[i][diff]表示下标i，差为diff的最长等差数列的长度

        