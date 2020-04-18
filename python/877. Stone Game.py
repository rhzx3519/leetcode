#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def stoneGame(self, piles):
        """
        :type piles: List[int]
        :rtype: bool
        """
        n = len(piles)
        dp = [[[0, 0] for _ in range(n)] for _ in range(n)]
        for i in range(n):
            dp[i][i] = [piles[i], 0]

        for l in range(1, n):
            for i in range(0, n-l):
                j = l+i
                
                left = piles[i] + dp[i+1][j][1]
                right = piles[j] + dp[i][j-1][1]
                if left>right:
                    dp[i][j][0] = left
                    dp[i][j][1] = dp[i+1][j][0]
                else:
                    dp[i][j][0] = right
                    dp[i][j][1] = dp[i][j-1][0]
                # print i, j, dp[i][j], left, right
        a, b = dp[0][n-1]
        return a>b

# 思路： dp[i][j][0]表示对i到j的石子堆先手的最大分数,
#       dp[i][j][1]表示对i到j的石子堆后手的最大分数

if __name__ == '__main__':
    piles = [3,7,2,3]
    su = Solution()
    su.stoneGame(piles)                    
# 0 1
# 1 2
# 2 3
# 0 2
# 1 3
# 0 3