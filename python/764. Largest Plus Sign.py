#!usr/bin/env python  
#-*- coding:utf-8 _*-  


class Solution(object):
    def orderOfLargestPlusSign(self, N, mines):
        """
        :type N: int
        :type mines: List[List[int]]
        :rtype: int
        """
        grid = [[N]*N for _ in range(N)]
        for x in mines:
            grid[x[0]][x[1]] = 0

        for k in range(N):
            leftOne = rightOne = upOne = downOne = 0
            i = 0
            j = N-1
            while i<N and j>=0:
                leftOne = 0 if grid[k][i]==0 else leftOne+1
                grid[k][i] = min(grid[k][i], leftOne)

                rightOne = 0 if grid[k][j]==0 else rightOne+1
                grid[k][j] = min(grid[k][j], rightOne)

                upOne = 0 if grid[i][k]==0 else upOne+1
                grid[i][k] = min(grid[i][k], upOne)

                downOne = 0 if grid[j][k]==0 else downOne+1
                grid[j][k] = min(grid[j][k], downOne)                

                i += 1
                j -= 1

        res = 0
        for i in range(N):
            for j in range(N):
                res = max(res, grid[i][j])
        return res


## 二次循环遍历，求左右上下连续1的数量