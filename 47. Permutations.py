class Solution(object):
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        self.dfs(res, nums, 0, len(nums) - 1)
        return res
        
    def dfs(self, res, nums, begin, end):
        if begin == end:
            res.append(nums[:])
            return
        i = begin
        while i <= end:
            nums[i], nums[begin] = nums[begin], nums[i]
            self.dfs(res, nums, begin + 1, end)
            nums[i], nums[begin] = nums[begin], nums[i]
            i += 1
    
        