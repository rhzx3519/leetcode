class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        i = 0
        ln = len(nums)
        while i < ln:
            if nums[i] == i + 1:
                i += 1
            else:
                if nums[i] >= 1 and nums[i] <= ln and nums[i] != nums[nums[i] - 1]:
                    nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]
                else:
                    i += 1
        
        i = 0
        while i < ln:
            if nums[i] != i + 1:
                return i+1
            i += 1
        return ln + 1
                    
s = Solution()
s.firstMissingPositive([3,4,-1,1])            