class Solution(object):
    def removeStones(self, stones):
        """
        :type stones: List[List[int]]
        :rtype: int
        """

        class UFS(object):
            def __init__(self, N):
                self.p = range(N)
            
            def find(self, x):
                if self.p[x] != x:
                    self.p[x] = self.find(self.p[x])
                return self.p[x]

            def union(self, x, y):
                px = self.find(x)
                py = self.find(y)
                self.p[py] = px
            
        ufs = UFS(20000)
        for x, y in stones:
            ufs.union(x, y + 10000)
        return len(stones) - len({ufs.find(x) for x, y in stones})

# 思路：并查集, 处于同行 或者同列的点属于同一个集合，列+10000来将二维的行列放到一个一维数组中
# time: O(NlogN), space: O(N)