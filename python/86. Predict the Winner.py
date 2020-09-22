class Solution(object):
    def PredictTheWinner(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        n = len(nums)
        dp = [[0]*n for _ in range(n)]
        for i in range(n):
            dp[i][i] = nums[i]
        
        for end in range(1, n): # 遍历方向有点怪异
            i = 0
            j = end
            while j < n:
                dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])
                i += 1
                j += 1
        return dp[0][n-1] >= 0


class Solution(object):
    def PredictTheWinner(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        mem = {}
        def dp(l, r):
            if l>=r:
                return nums[l]
            if (l, r) in mem:
                return mem[(l, r)]
            ans = max(nums[l] - dp(l+1, r), nums[r] - dp(l, r-1))
            mem[(l, r)] = ans
            return ans
        return dp(0, len(nums)-1) >= 0