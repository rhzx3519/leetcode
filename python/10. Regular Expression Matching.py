#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# 暴力递归
class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        if not p: return not s
        first_match = s and (s[0]==p[0] or p[0]=='.')
        if len(p)>=2 and p[1]=='*':
            # self.isMatch(s, p[2:]) 匹配0个前面字符
            # (first_match and self.isMatch(s[1:], p))  匹配1个或者多个前面字符
            return self.isMatch(s, p[2:]) or bool(first_match and self.isMatch(s[1:], p)) 
        else:
            return bool(first_match and self.isMatch(s[1:], p[1:]))


# 带备忘录的递归
class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        memo = {}

        def dp(i, j):
            if (i, j) in memo: return memo[(i, j)]
            if j==len(p): return i==len(s)
            first = i<len(s) and p[j] in (s[i], '.')
            if j < len(p)-1 and p[j+1]=='*':
                ans = dp(i, j+2) or (first and dp(i+1, j))
            else:
                ans = first and dp(i+1, j+1)
            memo[(i, j)] = ans
            return ans

        return dp(0, 0)
            

if __name__ == '__main__':
    s = "ab"
    p = ".*c"
    su = Solution()
    print su.isMatch(s, p)        
