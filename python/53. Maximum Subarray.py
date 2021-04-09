class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        #Kadane's algorithm
        ans = cur = None
        for x in nums:
            cur = x + max(cur, 0)
            ans = max(cur, ans)
        return ans

# 动态规划： 求和，然后判断和是否小于0，因为只要前面的和小于0，那么后面的数加上前面的和就一定比自身小，所以又重新求和，并和之前的最大子序和比较，取最大值。