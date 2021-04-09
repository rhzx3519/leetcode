class Solution(object):
    def numSubseq(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        Mod = 10**9 + 7
        nums.sort()
        l, r = 0, len(nums)-1
        ans = 0
        while l <= r:
            if nums[l] + nums[r] > target:
                r -= 1
            else:
                ans = (ans + (1<<(r-l)))%Mod
                l += 1
        return ans
