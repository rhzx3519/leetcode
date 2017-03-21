class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        ln = len(nums)
        nums.sort()
        i,j,k = 0,0,0
        while i < ln:
            l, h = i+1, ln-1
            while l < h:
                val = nums[i] + nums[l] + nums[h]                
                if val > 0:
                    h -= 1
                    continue
                if val < 0:
                    l += 1
                    continue

                if val == 0:
                    res.append([nums[i], nums[l], nums[h]])

                while l < h and nums[l] == nums[l + 1]:
                    l += 1
                l += 1

            while i < ln - 1 and nums[i] == nums[i+1]:
                i += 1
            i += 1
        
        return res

s = Solution()
print s.threeSum([0,0,0,0])