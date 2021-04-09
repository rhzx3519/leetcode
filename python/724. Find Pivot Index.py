class Solution(object):
    def pivotIndex(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        s = sum(nums)
        t = 0
        for i in range(len(nums)):
            if t == s-t-nums[i]:
                return i
            t += nums[i]

        return -1