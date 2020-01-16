# -*- coding: utf-8 -*-


class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        sc, si, sj = [], [], []
        ls, lp = len(s), len(p)
        i, j = 0, 0
        while i < ls or j < lp:
            # match star
            print i, j
            if j + 1 < lp and p[j+1] == '*':
                sc.append(p[j]) # 入栈等会儿比较
                si.append(i)
                j += 2
                sj.append(j)
            elif j < lp and i < ls and (s[i] == p[j] or p[j] == '.'): # 正常情况
                i += 1
                j += 1
            elif sc:
                # turn back and try again
                k = si[-1]
                if k < ls and (s[k] == sc[-1] or sc[-1] == '.'):
                    k += 1
                    i = si[-1] = k
                    j = sj[-1]
                else:
                    sc.pop()
                    si.pop()
                    sj.pop()
            else: # miss match
                break
        
        return i == ls and j == lp
                
s = Solution()
print s.isMatch('aaca', "ab*a*c*a")   


