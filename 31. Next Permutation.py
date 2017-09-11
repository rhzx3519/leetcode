class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        i = len(nums) - 1
        while i >= 0:
            j = len(nums) - 1
            while j > i:
                if nums[i] < nums[j]:
                    nums[i], nums[j] = nums[j], nums[i]
                    t = nums[i+1:]
                    t.sort()
                    k = 0
                    while k < len(t):
                        nums[k + i + 1] = t[k]
                        k += 1
                    return
                j -= 1
            i -= 1
            
        nums.sort()

s = Solution()
nums = [3,2,1]
s.nextPermutation(nums)
print(nums)
