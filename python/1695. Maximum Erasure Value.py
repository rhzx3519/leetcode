class Solution(object):
    def maximumUniqueSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        ans = 0
        l = s = 0
        mem = {}
        n = len(nums)
        for r in range(n):
            s += nums[r]
            
            while l < r and nums[r] in mem: 
                s -= nums[l]
                del mem[nums[l]]
                l += 1
            mem[nums[r]] = r
            ans = max(ans, s)
        return ans

# 思路：滑动窗口