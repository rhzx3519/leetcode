class Solution(object):
    def wiggleMaxLength(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        if n<2: return n
        dp = [[1, 1] for _ in range(n+1)]

        for i in range(1, n+1):
            for j in range(1, i):
                if nums[j-1] < nums[i-1]:
                    dp[i][0] = max(dp[i][0], dp[j][1] + 1)
                elif nums[j-1] > nums[i-1]:
                    dp[i][1] = max(dp[i][1], dp[j][0] + 1)
        
        return max(dp[n][0], dp[n][1])
# 思路： dp[i][0]表示i结尾的递增子序列长度, dp[i][1]表示i结尾的递减子序列长度

        