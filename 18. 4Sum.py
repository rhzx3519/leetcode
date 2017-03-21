class Solution(object):
    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        res = []
        nums.sort()
        ln = len(nums)
        i = 0
        while i < ln:
            j = i + 1
            while j < ln:
                l, h = j + 1, ln - 1
                while l < h:
                    val = nums[i] + nums[j] + nums[l] + nums[h]
                    if val < target:
                        l += 1
                        continue
                    if val > target:
                        h -= 1
                        continue
                    
                    res.append([nums[i], nums[j], nums[l], nums[h]])
                    while l < h-1 and nums[l] == nums[l+1]:
                        l += 1
                    l += 1

                while j < ln - 1 and nums[j] == nums[j+1]:
                    j += 1
                j += 1
            while i < ln - 1 and nums[i] == nums[i + 1]:
                i += 1
            i += 1
        return res

s = Solution()
print s.fourSum([-3,-1,0,2,4,5], 0)