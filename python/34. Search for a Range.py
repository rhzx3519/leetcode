class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        l, r = 0, len(nums) - 1
        while l <= r:
            mid = l+r>>1
            if nums[mid] == target:
                i = j = mid
                while i-1 >= 0 and nums[i-1] == nums[i]:
                    i -= 1
                while j+1 < len(nums) and nums[j+1] == nums[j]:
                    j += 1
                return [i, j]
            elif nums[mid] > target:
                r = mid - 1
            else:
                l = mid + 1
                
        return [-1, -1]