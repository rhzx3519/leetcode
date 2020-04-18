class Solution(object):
    def subsetsWithDup(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        n = len(nums)
        res = [[]]
        nums.sort()
        last = 0
        for i in range(n):
            l = 0
            if i>0 and nums[i-1]==nums[i]:
                l = last
            r = len(res)

            for j in range(l, r):
                res.append(res[j] + [nums[i]])
            last = r
                
        return res

# 思路: last记录上一轮加入到结果中子集，如果这轮的数字是重复的，则从last位置开始生成子集