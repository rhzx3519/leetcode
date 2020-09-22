
#!usr/bin/env python  
#-*- coding:utf-8 _*- 

class Solution(object):
    def combinationSum3(self, n):
        """
        :type k: int
        :type n: int
        :rtype: List[List[int]]
        """
        res = []
        track = []

        def dfs(idx):
            if idx==n+1:
                res.append(track[:])
                print track
                return 

            for i in range(idx, n+1):
                track.append(i)
                dfs(i+1) 
                track.pop()

        dfs(1)
        return res

if __name__ == '__main__':
    n = 3
    su = Solution()
    su.combinationSum3(n)