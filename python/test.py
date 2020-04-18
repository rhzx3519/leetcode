#!usr/bin/env python  
#-*- coding:utf-8 _*-  
import functools

# -*- coding:utf-8 -*-
class Solution:
    # s 源字符串
    def replaceSpace(self, s):
        # write code here
        mem = {}
        cnt = 0
        for ch in s:
            if ch==' ':
                cnt += 1
        t = ['']*(len(s) + cnt*2)
        i, j = len(s)-1, len(t)-1
        while i>=0 and j>=0:
            if s[i]==' ':
                t[j] = '0'
                j -= 1
                t[j] = '2'
                j -= 1
                t[j] = '%'
            else:
                t[j] = s[i]
            j -= 1
            i -= 1

        return ''.join(t)

if __name__ == '__main__':
    su = Solution()
    print su.replaceSpace('hello world')                    