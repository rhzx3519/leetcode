class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        k = 0
        for i in range(n):
            if nums[i]!=0:
                nums[k] = nums[i]
                k += 1
        while k < n:
            nums[k] = 0 
            k += 1

# 思路：k记录数组中第一个0的下标，碰到nums[i]!=0, 则和nums[k]交换，然后k+1
# time: O(N), space: O(1)