#!usr/bin/env python  
#-*- coding:utf-8 _*-  

import collections

class Solution(object):
    def patternMatching(self, pattern, value):
        """
        :type pattern: str
        :type value: str
        :rtype: bool
        """
        if not pattern and not value:
            return True

        count = collections.Counter(pattern)
        n = len(value)
        i = 0
        while i <= n:
            if i*count['a'] > n:
                break

            tmp = n - i*count['a']
            if count['b'] != 0 and tmp%count['b'] != 0:
                i += 1
                continue
            j = tmp/count['b'] if count['b'] != 0 else 0

            # a对应的字符传长度为i,b对应的字符串长度为j
            if self.match(pattern, value, i, j, count['a'], count['b']):

                return True
            i += 1

        return False

    def match(self, pattern, value, len_a, len_b, num_a, num_b):
        mem = {}
        i = 0
        n = len(value)

        for p in pattern:
            if p=='a':
                tmp = value[i:i+len_a]
                if (mem.has_key('a') and mem['a']!=tmp) or (mem.has_key('b') and mem['b']==tmp):
                    return False
                mem['a'] = tmp
                i += len_a
                num_a -= 1                    
            elif p=='b':
                tmp = value[i:i+len_b]
                if (mem.has_key('b') and mem['b']!=tmp) or (mem.has_key('a') and mem['a']==tmp):
                    return False
                mem['b'] = tmp
                i += len_b
                num_b -= 1
            else:
                return False

            if num_a<0 or num_b<0:
                return False

        return i==len(value)

# 思路： 就硬解呗
# time: O(M*N), space:O(M)
if __name__ == '__main__':
    pattern = 'aa'
    value = 'XXXX'
    su = Solution()
    print su.patternMatching(pattern, value) 

            