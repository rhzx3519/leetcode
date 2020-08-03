# -*- coding: utf-8 -*-

class Solution(object):
    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        lo, hi = 0, len(nums)-1
        while lo < hi:
            mid = lo+(hi-lo)/2
            if nums[mid] > nums[hi]:
                lo = mid+1
            elif nums[mid] < nums[hi]:
                hi = mid
            else:
                hi -= 1
        return nums[lo]

# 和 I 的做法类似, 都是二分法, 每次进入无序的那部分找出最小值
# 但是由于有重复值的情况, 需要加入 mid 元素等于 hi 元素的情况
# 此时应该将 hi 减 1 防止重复数字是最小元素