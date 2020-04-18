#!usr/bin/env python  
#-*- coding:utf-8 _*-  
from collections import Counter
class Solution(object):
    def maxRepOpt1(self, text):
        """
        :type text: str
        :rtype: int
        """
        mem = Counter(text)
        i = 0
        n = len(text)
        res = 0
        while i < n:
            ch = text[i]
            j = i+1
            count = 1
            idx = i
            fake = True
            while j<n and (ch==text[j] or fake):
                if ch!=text[j]:
                    fake = False
                    idx = j-1
                j += 1
                count += 1
            i = idx
            i += 1
            count = min(count, mem[ch])
            res = max(res, count)
        return res
        
if __name__ == '__main__':
    text = "aaabbaaa"
    su = Solution()
    print su.maxRepOpt1(text)        