class Solution(object):
    def canPlaceFlowers(self, flowerbed, n):
        """
        :type flowerbed: List[int]
        :type n: int
        :rtype: bool
        """
        cnt = 0
        m = len(flowerbed)
        for i in range(m):
            if flowerbed[i]==1: continue
            l = (i==0) or (flowerbed[i-1]==0)
            r = (i==m-1) or (flowerbed[i+1]==0)
            if l and r:
                flowerbed[i] = 1
                n -= 1
        return n<=0
        