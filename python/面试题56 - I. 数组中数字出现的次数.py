#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def singleNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = 0
        for num in nums:
            n^=num
        t = n&(-n)
        r1 = r2 = 0
        for num in nums:
            if num&t:
                r1 ^= num
            else:
                r2 ^= num
        return [r1, r2]

# 思路：n&(-n)返回从右向左数第一个1和右边0构成的十进制数
    