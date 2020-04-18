class Solution(object):
    def findNumberOfLIS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        dp = [1]*(n+1)
        cnt = [1]*(n+1)
        cnt[0] = dp[0] = 0
        for i in range(1, n+1):
            for j in range(1, i):
                if nums[j-1] < nums[i-1]:
                    if dp[i]<dp[j]+1:
                        dp[i] = dp[j]+1
                        cnt[i] = cnt[j]
                    elif dp[i]==dp[j]+1:
                        cnt[i] += cnt[j]

        res = 0
        max_val = max(dp)
        for i in range(n+1):
            if dp[i]==max_val:
                res += cnt[i]
        return res
        