class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = [[]]
        for i in range(len(nums)-1, -1, -1):
            for subset in res[:]:
                res.append(subset + [nums[i]])
        return res

# 思路：从后向前遍历，遇到一个数就把所有子集加上该数形成新的子集