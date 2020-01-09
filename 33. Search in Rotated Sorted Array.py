class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        if not nums:
            return -1
        ls = len(nums)
        i = 0
        while i < ls - 1:
            if nums[i] > nums[i+1]:
                break
            i += 1
        l, r = 0, ls - 1
        if target >= nums[0] and target <= nums[i]:
            r = i
        else:
            l = i + 1
            
        while l <= r:
            mid = (l + r) >> 1
            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                r = mid - 1
            else:
                l = mid + 1
        
        return -1