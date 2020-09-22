class Solution(object):
    def longestSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if all(nums):
            return len(nums) - 1
        n = len(nums)
        l = 0
        cnt = 0
        max_val = 0
        for r, bit in enumerate(nums):
            if bit==0:
                cnt += 1
            while cnt > 1:
                if nums[l]==0:
                    cnt -= 1
                l += 1
            max_val = max(max_val, r - l)
        return max_val