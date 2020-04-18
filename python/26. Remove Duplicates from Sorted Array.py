class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        i = j = k = 0
        ln = len(nums)
        while i < ln:
            nums[k] = nums[i]
            k += 1
            j = i + 1
            while j < ln and nums[i] == nums[j]:
                j += 1
            i = j
        return k
        