#!usr/bin/env python  
#-*- coding:utf-8 _*-  

class Solution(object):
    def connectTwoGroups(self, cost):
        """
        :type cost: List[List[int]]
        :rtype: int
        """
        m, n = len(cost), len(cost[0])
        costMatrix = [[0]*(1<<n) for _ in range(m)]
        for i in range(m):
            for k in range(1<<n):
                s = 0
                for j in range(n):
                    if k&(1<<j) > 0:
                        s += cost[i][j]
                costMatrix[i][k] = s
        # print costMatrix
        dp = [[float('inf')]*(1<<n) for _ in range(m)]
        dp[0] = costMatrix[0]
        for i in range(1, m):
            for j in range(1, 1<<n):
                for k in range(1, 1<<n):
                    dp[i][j|k] = min(dp[i][j|k], dp[i-1][k] + costMatrix[i][j])
        return dp[m-1][(1<<n)-1]

if __name__ == '__main__':
    cost = [[15, 96], [36, 2]]
    su = Solution()
    print su.connectTwoGroups(cost)

# costMatrix[i][j] 表示第 i 行选取状况为 j 时该行被选取得元素总和