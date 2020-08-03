class Solution(object):
    def findRepeatNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        for i in range(n):
            while i != nums[i]:
                t = nums[i]
                if nums[i] == nums[t]:
                    return t
                nums[i], nums[t] = nums[t], nums[i]
        return -1
        
# time: O(N), space: O(1)