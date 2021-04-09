class Solution(object):
    def movesToMakeZigzag(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        odd = even = 0
        for i in range(n):
            if i&1: # i-1, i+1是偶数
                d1 = nums[i] - nums[i-1] + 1 if nums[i] >= nums[i-1] else 0
                d2 = nums[i] - nums[i+1] + 1 if i < n-1 and nums[i] >= nums[i+1] else 0
                odd += max(d1, d2)
            else:
                d1 = nums[i] - nums[i-1] + 1 if i > 0 and nums[i] >= nums[i-1] else 0
                d2 = nums[i] - nums[i+1] + 1 if i < n-1 and nums[i] >= nums[i+1] else 0
                even += max(d1, d2)
        return min(odd, even)

# 思路：下标分成奇数或者偶数的情况分别计算
