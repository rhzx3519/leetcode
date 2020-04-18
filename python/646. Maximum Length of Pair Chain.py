class Solution(object):
    def findLongestChain(self, pairs):
        """
        :type pairs: List[List[int]]
        :rtype: int
        """
        if not pairs: return 0
        pairs.sort(key=lambda k: k[1])
        res = 1
        end = pairs[0][1]
        for i in range(1, len(pairs)):
            if pairs[i][0] > end:
                end = pairs[i][1]
                res += 1
        return res; 

# 求最大非重叠区间