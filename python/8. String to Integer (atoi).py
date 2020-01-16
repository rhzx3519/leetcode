#!/usr/bin/env python
# -*- coding: utf-8 -*-

class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        i = 0
        while i < len(str) and str[i] == ' ':
            i += 1
        if i >= len(str):
            return 0
        sign = 1
        if str[i] == '+':
            i += 1
        elif str[i] == '-':
            i += 1
            sign = 0
        if not str[i:]:
            return 0
            
        start = i
        while i < len(str):
            if str[i] < '0' or str[i] >'9':
                break
            i += 1
        if i == start:
            return 0
        # 判断是否超过最大或者最小边界
        res = int(str[start:i])
        if sign:
            res = min(pow(2, 31) - 1, res)
        else:
            res = max(-pow(2, 31), -res)
        
        return res

