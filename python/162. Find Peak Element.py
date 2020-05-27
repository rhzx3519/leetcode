class Solution(object):
    def findPeakElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        l = 0
        r = n-1
        while l<r:
            m1 = l + (r-l)/2
            m2 = m1+1
            if nums[m1] < nums[m2]:
                l = m2
            else:
                r = m1
        return l