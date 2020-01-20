# -*- coding: utf-8 -*-

class Solution(object):
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        if len(nums)<2:
            return True
        i = pos = new_pos = 0
        while True:
            new_pos = pos    
            while i <= pos:
                new_pos = max(new_pos, i + nums[i])
                i += 1

            if new_pos==pos:
                return False

            pos = new_pos
            if pos >= len(nums) - 1:
                return True

        return False