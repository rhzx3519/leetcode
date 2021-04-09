#!usr/bin/env python  
#-*- coding:utf-8 _*- 

import collections

class Solution(object):
    def removeDuplicateLetters(self, s):
        """
        :type s: str
        :rtype: str
        """
        st = []
        count = collections.Counter(s)
        for c in s:
            if c not in st:
                while st and c < st[-1] and count[st[-1]] > 0:
                    # count[st[-1]] > 0 说明后面还有st[-1]
                    st.pop()
                st.append(c)
            count[c] -= 1
        return ''.join(st)

if __name__ == '__main__':
    s = "bcabc"
    su = Solution()
    print su.removeDuplicateLetters(s)