#!/usr/bin/python
# -*- coding: UTF-8 -*-

import collections

class Solution(object):
    def characterReplacement(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: int
        """
        l = r= 0
        max_val = 0
        count = collections.defaultdict(int)
        res = 0
        for i in range(len(s)):
            count[s[i]] += 1
            max_val = max(max_val, count[s[i]])
            while i-l+1-max_val > k:
                count[s[l]] -= 1
                l += 1
            res = max(res, i-l+1)
        return res

# 滑动窗口，需要让窗口中的重复元素尽可能的多，当窗口中除最多元素之外的剩余元素数量大于k时，窗口左边界右移

