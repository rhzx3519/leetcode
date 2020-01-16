class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        i = j = k = 0
        while i < len(nums):
            while i < len(nums) and nums[i] == val:
                i += 1
            if i < len(nums):
                nums[k] = nums[i]
                k += 1
            i += 1    
        return k
        