class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        output = [1] * n
        product = 1
        for i in range(n-1, -1, -1):
            output[i] *= product
            product *= nums[i]
        
        product = 1
        for i in range(n):
            output[i] *= product
            product *= nums[i]
        return output
            