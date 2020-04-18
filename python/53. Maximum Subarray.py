class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        s = 0
        max_val = -1<<31
        for num in nums:
            if s < 0: s = 0
            s += num
            max_val = max(max_val, s)
        return max_val

# 动态规划： 求和，然后判断和是否小于0，因为只要前面的和小于0，那么后面的数加上前面的和就一定比自身小，所以又重新求和，并和之前的最大子序和比较，取最大值。