class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        tmp = [0]*n
        for i in range(len(nums)):
            if 0<nums[i]<=n:
                tmp[nums[i] - 1] = nums[i]
                
        for i in range(n):
            if tmp[i] != i+1:
                return i+1
        return n+1

# 遍历一次数组把大于等于1的和小于数组大小的值放到原数组对应位置，然后再遍历一次数组查当前下标是否和值对应，
# 如果不对应那这个下标就是答案，否则遍历完都没出现那么答案就是数组长度加1。        