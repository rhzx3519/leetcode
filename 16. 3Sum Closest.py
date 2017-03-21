class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        res = target
        min_diff = 1<<29
        i, ln = 0, len(nums)
        nums.sort()
        while i < ln:
            l, h = i+1, ln-1
            while l < h:
                val = nums[i] + nums[l] + nums[h]
                if target == val:
                    res = val
                    return res

                if abs(target - val) < min_diff:
                    min_diff = abs(target - val)        
                    res = val
                
                if val < target:
                    l += 1
                    continue
                if val > target:
                    h -= 1
                    continue
                while l < ln - 1 and nums[l] == nums[l+1]:
                    l += 1
                
                l += 1
                
            while i < ln-1 and nums[i] == nums[i+1]:
                i += 1
            i += 1
        return res
        
s = Solution()
print s.threeSumClosest([-1,2,1,-4], 1)        