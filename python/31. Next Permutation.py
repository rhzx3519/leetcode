class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        for i in range(n-1, -1, -1):
            for j in range(n-1, i, -1):
                if nums[j] > nums[i]:
                    nums[i], nums[j] = nums[j], nums[i]
                    nums[i+1:] = sorted(nums[i+1:])
                    return
        nums.sort()

s = Solution()
nums = [3,2,1]
s.nextPermutation(nums)
print(nums)

