# -*- coding: utf-8 -*-

class Solution(object):
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        max_dis = 0 # 记录下标为i时，到达的最大距离
        for i in range(len(nums)):
            if i > max_dis:
                return False
            max_dis = max(max_dis, i + nums[i])
            if max_dis >= len(nums)-1:
                return True
        return True

# 思路：贪心