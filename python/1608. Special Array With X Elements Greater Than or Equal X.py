class Solution(object):
    def specialArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums.sort()
        n = len(nums)
        if nums[0] >= n: return n
        for i in range(1, n+1):
            if nums[n-1-i] < i and nums[n-i] >= i:
                return i
        return -1
