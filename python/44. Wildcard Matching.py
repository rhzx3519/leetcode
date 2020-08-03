#!usr/bin/env python  
#-*- coding:utf-8 _*-  


class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        ls = len(s)
        lp = len(p)
        dp = [[0]*(lp+1) for _ in range(ls+1)]
        dp[0][0] = 1
        for i in range(1, lp+1):
            dp[0][i] = dp[0][i-1] and p[i-1]=='*'

        for i in range(1, ls+1):
            for j in range(1, lp+1):
                if p[j-1]=='*':
                    dp[i][j] = dp[i][j-1] or dp[i-1][j] or dp[i-1][j-1]
                elif p[j-1]==s[i-1] or p[j-1]=='?':
                    dp[i][j] = dp[i-1][j-1]
                else:
                    dp[i][j] = 0
        return bool(dp[ls][lp])

# /*
# 定义dp[s.size()+1][p.size()+1].这样可以处理串为空的情况。
# dp[i][j]表示s[i-1]和p[j-1]（包括这两点）以及之前的字符串能否匹配。

# 初始dp[0][0]为1。因为两个空串匹配。
# 初始化dp[0][i]。表示当s为空串时，p是否全是*。
# dp[0][i]=dp[0][i-1] and p[i-1]=='*';

# 两重循环中。

# 当p[j-1]=*时。dp[i][j]在一下3种情况下为1。
# 下面用 i,j 和表示s和p所在的点，利于说明。
# （1）s在i点之前的字符串与p在j点之前的字符串匹配。
# （2）s（包括i点）和p在j点之前的字符串匹配
# （3）s在i点之前的点和p（包括p点）之前的字符串匹配。
# 因为 * 是范围匹配的。

# 当p[j]=?或p[j]=s[i]时。
# 表示这两个点匹配。所以s[i]和p[j]（包括这两点）的字符串能否匹配
# 取决于s[i-1]和p[j-1]（包括这两点）之前的字符串能否匹配。

# 当p[j]！=s[i]时。
# s[i]和p[j]（包括这两点）的字符串肯定不能匹配。

# 如有错误，还请指出！谢谢。
# */        

    def isMatch1(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        mem = {}
        def dp(i, j):
            if (i, j) in mem:
                return mem[(i, j)]

            if j==len(p):
                return i==len(s)
            first = i<len(s) and p[j] in (s[i], '?')
            if p[j]=='*':
                ans = (i<len(s)-1 and dp(i+1, j)) or dp(i, j+1) or dp(i+1, j+1)
            else:
                ans = first and dp(i+1, j+1)
            mem[(i, j)] = ans
            return ans

        return dp(0, 0)

if __name__ == '__main__':
    s = ""
    p = "*"
    su = Solution()
    print su.isMatch1(s, p)





















