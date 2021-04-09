class Solution(object):
    def minDifference(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        min_val = float('inf')
        n = len(nums)
        if n<=4: return 0
        nums.sort()
        # print nums
        N = 4
        for i in range(N):
            min_val = min(min_val, abs(nums[n-N+i] - nums[i]))
        return min_val

# 思路，排序，考虑去掉首尾三个元素之后，剩余数组中最大值与最小值的差值