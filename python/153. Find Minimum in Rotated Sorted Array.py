class Solution(object):
    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        l, r = 0, len(nums)-1
        while l < r:
            m = (l+r)>>1
            k = nums[m]
            if k < nums[r]: # m->r递增
                r = m
            else:
                l = m + 1
        return nums[l]